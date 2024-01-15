package controllers

import (
	"github.com/go-home-admin/home/bootstrap/utils"
	"testing"
)

func TestFiles_GetTitleList(t *testing.T) {
	f := &Files{}
	gotGot, gotCount := f.GetTitleList(1, 11)

	utils.Dump(gotGot, gotCount)
}
