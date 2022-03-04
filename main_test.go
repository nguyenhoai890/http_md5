package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func TestGetMD5HashResponse(t *testing.T) {
	// Since there is a limit lines of code. We can't test the returned hash without mock interface. So there is no testing same response data with same hash.
	type testCase struct {
		url         string
		hasErr      bool
		description string
	}
	testCases := []testCase{
		{
			url:         "",
			hasErr:      true,
			description: "blank url will return error",
		},
		{
			url:         "thisisnotaurl",
			hasErr:      true,
			description: "invalid url will return error",
		},
		{
			url:         fmt.Sprintf("%s.com", uuid.New().String()),
			hasErr:      true,
			description: "invalid website will return error",
		},
		{
			url:         "https://adjust.com",
			description: "valid url should not return error",
		},
		{
			url:         "adjust.com",
			description: "url without Scheme(http|https) should not return error",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := getMD5HashResponse(tc.url)
			if tc.hasErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestParseParams(t *testing.T) {
	type testCase struct {
		iUrls       []string
		iParallel   *int
		eUrls       []string
		eParallel   int
		description string
	}
	parallel11 := 11
	parallel0 := 0
	testCases := []testCase{
		{
			iUrls:       []string{"google.com"},
			iParallel:   &parallel11,
			eUrls:       []string{"google.com"},
			eParallel:   parallel11,
			description: "happy case",
		},
		{
			iUrls:       []string{"google.com", "adjust.com", "yahoo.com"},
			iParallel:   &parallel11,
			eUrls:       []string{"google.com", "adjust.com", "yahoo.com"},
			eParallel:   parallel11,
			description: "happy case with multiple domains",
		},
		{
			iUrls:       []string{},
			iParallel:   &parallel0,
			eUrls:       []string{},
			eParallel:   defaultParallel,
			description: fmt.Sprintf("invalid parallel param value, %v, should return %v, default value", parallel0, defaultParallel),
		},
		{
			iUrls:       []string{},
			iParallel:   nil,
			eUrls:       []string{},
			eParallel:   defaultParallel,
			description: fmt.Sprintf("unset parallel param should return %v, default value", defaultParallel),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			os.Args = []string{"test"}
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if tc.iParallel != nil {
				os.Args = append(os.Args, fmt.Sprintf("-%s", FlagParallel), strconv.Itoa(*tc.iParallel))
			}
			for _, url := range tc.iUrls {
				os.Args = append(os.Args, url)
			}
			parallel, urls := parseParams()
			require.Equal(t, tc.eUrls, urls)
			require.Equal(t, tc.eParallel, parallel)
			flag.Parse()
		})
	}
}

func TestRacePrintMD5HashResponses(t *testing.T) {
	numberOfUrls := 1000
	parallel := 10
	urls := make([]string, 0, numberOfUrls)
	for i :=0; i < numberOfUrls; i++ {
		fakeUrl := fmt.Sprintf("%s.com", uuid.New().String())
		urls = append(urls, fakeUrl)
	}
	printMD5HashResponses(urls, parallel)
}