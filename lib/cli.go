package lib

import "github.com/fatih/color"

// Catchy : bold and cyan text
var Catchy = color.New(color.Bold, color.FgCyan).PrintfFunc()

// Notice : bold and blue text
var Notice = color.New(color.Bold, color.FgBlue).PrintfFunc()

// Warning : yellow text
var Warning = color.New(color.FgYellow).PrintfFunc()

// Error : bold and red text
var Error = color.New(color.Bold, color.FgRed).PrintfFunc()
