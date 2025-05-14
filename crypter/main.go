package main

import (
	"context"
	"fmt"
	"github.com/mhthrh/common_pkg/util/cryptox"
	"github.com/mhthrh/common_pkg/util/file/directory"
	"github.com/mhthrh/common_pkg/util/file/text"
	"github.com/mhthrh/crypto/model"
	"github.com/mhthrh/crypto/validity"
	"reflect"
)

func initAction() *model.Action {
	return &model.Action{
		Key: model.Value{
			ID:      1,
			Message: "Enter the secret key to encrypt/decrypt your data:",
			Value:   "",
		},
		Type: model.Value{
			ID:      2,
			Message: "Choose operation: (e)encrypt / (d)decrypt:",
			Value:   "",
		},
		From: model.Value{
			ID:      3,
			Message: "Source file path:",
			Value:   "",
		},
	}
}

var (
	cmd string
	c   chan model.Action
	e   chan error
)

func init() {
	c, e = make(chan model.Action), make(chan error)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Printf("version: %s\n", model.Ver)
	fmt.Println("Welcome to crypter — your secure companion for encryption and decryption.")

	go crypto(ctx, c, e)
	for {
		msg := initAction()
		if err := getInfo(msg); err != nil {
			fmt.Println(err)
			continue
		}

		c <- *msg
		if result := <-e; result != nil {
			fmt.Printf("enc/dec failed: %v\n", result)
		} else {
			fmt.Printf("enc/dec successfull")
		}
		fmt.Println("Would you like to continue with another file? (yes/no): ")
		_, err := fmt.Scanln(&cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if cmd == "y" {
			continue
		}
		cancel()
		break
	}

}

func getInfo(i interface{}) error {
	cmd = ""
	v := reflect.ValueOf(i).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Struct && field.CanSet() {
			id := field.FieldByName("ID")
			msgField := field.FieldByName("Message")
			valField := field.FieldByName("Value")
			if msgField.IsValid() && valField.IsValid() && valField.CanSet() && msgField.Kind() == reflect.String && valField.Kind() == reflect.String {
				fmt.Println(msgField)
				_, err := fmt.Scanln(&cmd)
				if err != nil {
					return err
				}
				cmd, err = validity.Validation(cmd, int(id.Int()))
				if err != nil {
					return err
				}
				valField.SetString(cmd)
			}
		}
	}
	return nil
}
func crypto(ctx context.Context, action chan model.Action, e chan error) {
	defer func() {
		close(action)
		close(e)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case a := <-action:

			fileInfo, err := directory.GetFileName(a.From.Value)
			if err != nil {
				e <- err
				continue
			}
			rFile := text.New("", a.From.Value, true)
			wFile := text.New("", fmt.Sprintf("%s/%s", fileInfo.Dir, fileInfo.Name), true)

			c, err := cryptox.New(a.Key.Value)
			if err != nil {
				e <- err
				continue
			}

			byt, err := rFile.Read()
			if err != nil {
				e <- err
				continue
			}

			if a.Type.Value == "e" {
				str, err := c.Encrypt(string(byt))
				if err != nil {
					e <- err
					continue
				}
				err = wFile.Write([]byte(str))
				if err != nil {
					e <- err
					continue
				}
				e <- nil
				continue
			}
			str, err := c.Decrypt(string(byt))
			if err != nil {
				e <- err
				continue
			}
			err = wFile.Write([]byte(str))
			if err != nil {
				e <- err
				continue
			}
			e <- nil
		}
	}
}
