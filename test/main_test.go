package test

import (
	"bytes"
	"github.com/fdev-ci/golang-plugin-sdk/log"
	"os/exec"
	"testing"
)

func Build(script string) error {

	cmd := exec.Command("sh", "-c",script)
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out

	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Info(err)
	}
	log.Info(out.String())
	return err

}
func TestBuild(t *testing.T) {
	Build("luotao")
}
