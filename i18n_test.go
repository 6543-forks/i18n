// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package i18n

import (
	"testing"
)

func Test_Tr(t *testing.T) {
	_ = SetMessage("en-US", "testdata/locale_en-US.ini")
	result := Tr("en-US", "NAME")
	if result != "English" {
		t.Errorf("expect 'English', got '%s'", result)
	}

	result = Tr("en-US", "section.NAME")
	if result != "Chinese" {
		t.Errorf("expect 'Chinese', got '%s'", result)
	}

	result = Tr("en-US", "SECONDS", 10)
	if result != "10 seconds" {
		t.Errorf("expect '10 seconds', got '%s'", result)
	}

	result = Tr("en-US", ".BAD.NAME")
	if result != "Bad Name" {
		t.Errorf("expect 'Bad Name', got '%s'", result)
	}

	result = Tr("en-US", "section.COMMENT_SYMBOL")
	if result != `test value <span style="color: ; background: ">more text</span>` {
		t.Errorf(`expect 'test value <span style="color: ; background: ">more text</span>', got '%s'`, result)
	}

	langs := ListLangs()
	if len(langs) != 1 {
		t.Errorf("expect '1', got '%d'", len(langs))
	} else if langs[0] != "en-US" {
		t.Errorf("expect 'en-US', got '%s'", langs[0])
	}

	Flush()

	langs = ListLangs()
	if len(langs) != 0 {
		t.Errorf("expect '0', got '%d'", len(langs))
	}
}

func Benchmark_Tr(b *testing.B) {
	_ = SetMessage("en-US", "testdata/locale_en-US.ini")
	for i := 0; i < b.N; i++ {
		Tr("en-US", "NAME")
	}
}

func Benchmark_TrWithSection(b *testing.B) {
	_ = SetMessage("en-US", "testdata/locale_en-US.ini")
	for i := 0; i < b.N; i++ {
		Tr("en-US", "section.NAME")
	}
}

func Benchmark_TrWithFormat(b *testing.B) {
	_ = SetMessage("en-US", "testdata/locale_en-US.ini")
	for i := 0; i < b.N; i++ {
		Tr("en-US", "SECONDS", 10)
	}
}
