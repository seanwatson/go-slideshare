//
// File: contact.go
// Date: 1/26/2014
//
// This file contains the response containers for API calls that return contacts.
//
// The MIT License (MIT)
//
// Copyright (c) 2014 Sean Watson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package slideshare

// A container for contacts represented in XML responses.
type Contact struct {
	Username      string `xml:"Username"`
	NumSlideshows uint32 `xml:"NumSlideshows"`
	NumComments   uint32 `xml:"NumComments"`
}

// Used when multiple contacts are returned in a single API call.
type Contacts struct {
	Values []Contact `xml:"Contact"`
}
