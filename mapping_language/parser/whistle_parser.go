// Code generated from ./parser/Whistle.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Whistle
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 336,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	5, 8, 80, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 85, 10, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 7, 9, 91, 10, 9, 12, 9, 14, 9, 94, 11, 9, 3, 9, 5, 9, 97, 10, 9, 3,
	9, 7, 9, 100, 10, 9, 12, 9, 14, 9, 103, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 113, 10, 10, 12, 10, 14, 10, 116, 11,
	10, 5, 10, 118, 10, 10, 3, 10, 3, 10, 5, 10, 122, 10, 10, 3, 10, 3, 10,
	5, 10, 126, 10, 10, 3, 11, 5, 11, 129, 10, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 5, 12, 135, 10, 12, 3, 12, 5, 12, 138, 10, 12, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19,
	163, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 169, 10, 19, 12, 19, 14,
	19, 172, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 178, 10, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 186, 10, 20, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 5, 24, 202, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 208, 10,
	24, 12, 24, 14, 24, 211, 11, 24, 5, 24, 213, 10, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 7, 24, 220, 10, 24, 12, 24, 14, 24, 223, 11, 24, 5, 24, 225,
	10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 231, 10, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 251, 10, 24, 12, 24, 14, 24,
	254, 11, 24, 3, 25, 3, 25, 5, 25, 258, 10, 25, 3, 25, 3, 25, 5, 25, 262,
	10, 25, 3, 25, 5, 25, 265, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 273, 10, 25, 5, 25, 275, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 285, 10, 26, 3, 27, 3, 27, 7, 27, 289,
	10, 27, 12, 27, 14, 27, 292, 11, 27, 3, 27, 5, 27, 295, 10, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 303, 10, 28, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 5, 29, 311, 10, 29, 3, 30, 3, 30, 7, 30, 315,
	10, 30, 12, 30, 14, 30, 318, 11, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 328, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 5,
	33, 334, 10, 33, 3, 33, 2, 3, 46, 34, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 2, 9, 3, 2, 25, 26, 3, 2, 27, 28, 3, 2, 29, 34, 3, 2, 36,
	37, 3, 3, 21, 21, 4, 2, 14, 14, 19, 19, 4, 2, 17, 18, 42, 42, 2, 368, 2,
	66, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 6, 70, 3, 2, 2, 2, 8, 72, 3, 2, 2, 2,
	10, 74, 3, 2, 2, 2, 12, 76, 3, 2, 2, 2, 14, 79, 3, 2, 2, 2, 16, 92, 3,
	2, 2, 2, 18, 106, 3, 2, 2, 2, 20, 128, 3, 2, 2, 2, 22, 132, 3, 2, 2, 2,
	24, 139, 3, 2, 2, 2, 26, 142, 3, 2, 2, 2, 28, 145, 3, 2, 2, 2, 30, 149,
	3, 2, 2, 2, 32, 153, 3, 2, 2, 2, 34, 157, 3, 2, 2, 2, 36, 160, 3, 2, 2,
	2, 38, 175, 3, 2, 2, 2, 40, 187, 3, 2, 2, 2, 42, 190, 3, 2, 2, 2, 44, 193,
	3, 2, 2, 2, 46, 230, 3, 2, 2, 2, 48, 274, 3, 2, 2, 2, 50, 284, 3, 2, 2,
	2, 52, 286, 3, 2, 2, 2, 54, 302, 3, 2, 2, 2, 56, 310, 3, 2, 2, 2, 58, 312,
	3, 2, 2, 2, 60, 319, 3, 2, 2, 2, 62, 327, 3, 2, 2, 2, 64, 333, 3, 2, 2,
	2, 66, 67, 9, 2, 2, 2, 67, 3, 3, 2, 2, 2, 68, 69, 9, 3, 2, 2, 69, 5, 3,
	2, 2, 2, 70, 71, 9, 4, 2, 2, 71, 7, 3, 2, 2, 2, 72, 73, 9, 5, 2, 2, 73,
	9, 3, 2, 2, 2, 74, 75, 7, 38, 2, 2, 75, 11, 3, 2, 2, 2, 76, 77, 7, 35,
	2, 2, 77, 13, 3, 2, 2, 2, 78, 80, 7, 27, 2, 2, 79, 78, 3, 2, 2, 2, 79,
	80, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 84, 7, 45, 2, 2, 82, 83, 7, 41,
	2, 2, 83, 85, 7, 45, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85,
	15, 3, 2, 2, 2, 86, 91, 5, 38, 20, 2, 87, 91, 5, 40, 21, 2, 88, 91, 5,
	18, 10, 2, 89, 91, 7, 21, 2, 2, 90, 86, 3, 2, 2, 2, 90, 87, 3, 2, 2, 2,
	90, 88, 3, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3,
	2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95,
	97, 5, 64, 33, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 101, 3, 2,
	2, 2, 98, 100, 7, 21, 2, 2, 99, 98, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101,
	99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 101, 3,
	2, 2, 2, 104, 105, 7, 2, 2, 3, 105, 17, 3, 2, 2, 2, 106, 107, 7, 39, 2,
	2, 107, 108, 7, 42, 2, 2, 108, 117, 7, 3, 2, 2, 109, 114, 5, 20, 11, 2,
	110, 111, 7, 4, 2, 2, 111, 113, 5, 20, 11, 2, 112, 110, 3, 2, 2, 2, 113,
	116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 118,
	3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 109, 3, 2, 2, 2, 117, 118, 3, 2,
	2, 2, 118, 119, 3, 2, 2, 2, 119, 121, 7, 5, 2, 2, 120, 122, 7, 21, 2, 2,
	121, 120, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	125, 5, 36, 19, 2, 124, 126, 7, 21, 2, 2, 125, 124, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 19, 3, 2, 2, 2, 127, 129, 7, 40, 2, 2, 128, 127, 3, 2,
	2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 7, 42, 2, 2,
	131, 21, 3, 2, 2, 2, 132, 134, 5, 24, 13, 2, 133, 135, 5, 26, 14, 2, 134,
	133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 138,
	7, 21, 2, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 23, 3, 2,
	2, 2, 139, 140, 5, 42, 22, 2, 140, 141, 5, 36, 19, 2, 141, 25, 3, 2, 2,
	2, 142, 143, 7, 13, 2, 2, 143, 144, 5, 36, 19, 2, 144, 27, 3, 2, 2, 2,
	145, 146, 7, 3, 2, 2, 146, 147, 5, 42, 22, 2, 147, 148, 7, 5, 2, 2, 148,
	29, 3, 2, 2, 2, 149, 150, 7, 22, 2, 2, 150, 151, 5, 44, 23, 2, 151, 152,
	7, 23, 2, 2, 152, 31, 3, 2, 2, 2, 153, 154, 7, 22, 2, 2, 154, 155, 7, 45,
	2, 2, 155, 156, 7, 23, 2, 2, 156, 33, 3, 2, 2, 2, 157, 158, 7, 22, 2, 2,
	158, 159, 7, 23, 2, 2, 159, 35, 3, 2, 2, 2, 160, 162, 7, 6, 2, 2, 161,
	163, 7, 21, 2, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 170,
	3, 2, 2, 2, 164, 169, 5, 38, 20, 2, 165, 169, 5, 40, 21, 2, 166, 169, 5,
	22, 12, 2, 167, 169, 7, 21, 2, 2, 168, 164, 3, 2, 2, 2, 168, 165, 3, 2,
	2, 2, 168, 166, 3, 2, 2, 2, 168, 167, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2,
	170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172,
	170, 3, 2, 2, 2, 173, 174, 7, 7, 2, 2, 174, 37, 3, 2, 2, 2, 175, 177, 5,
	50, 26, 2, 176, 178, 5, 28, 15, 2, 177, 176, 3, 2, 2, 2, 177, 178, 3, 2,
	2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 7, 8, 2, 2, 180, 185, 5, 46, 24,
	2, 181, 186, 7, 9, 2, 2, 182, 186, 5, 40, 21, 2, 183, 186, 7, 21, 2, 2,
	184, 186, 7, 2, 2, 3, 185, 181, 3, 2, 2, 2, 185, 182, 3, 2, 2, 2, 185,
	183, 3, 2, 2, 2, 185, 184, 3, 2, 2, 2, 186, 39, 3, 2, 2, 2, 187, 188, 7,
	46, 2, 2, 188, 189, 9, 6, 2, 2, 189, 41, 3, 2, 2, 2, 190, 191, 7, 11, 2,
	2, 191, 192, 5, 46, 24, 2, 192, 43, 3, 2, 2, 2, 193, 194, 7, 12, 2, 2,
	194, 195, 5, 46, 24, 2, 195, 45, 3, 2, 2, 2, 196, 197, 8, 24, 1, 2, 197,
	231, 5, 48, 25, 2, 198, 231, 5, 36, 19, 2, 199, 201, 7, 42, 2, 2, 200,
	202, 5, 34, 18, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203,
	3, 2, 2, 2, 203, 212, 7, 3, 2, 2, 204, 209, 5, 46, 24, 2, 205, 206, 7,
	4, 2, 2, 206, 208, 5, 46, 24, 2, 207, 205, 3, 2, 2, 2, 208, 211, 3, 2,
	2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 213, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 212, 204, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213,
	214, 3, 2, 2, 2, 214, 231, 7, 5, 2, 2, 215, 224, 7, 22, 2, 2, 216, 221,
	5, 46, 24, 2, 217, 218, 7, 4, 2, 2, 218, 220, 5, 46, 24, 2, 219, 217, 3,
	2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2,
	2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 216, 3, 2, 2, 2, 224,
	225, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 231, 7, 23, 2, 2, 227, 228,
	5, 12, 7, 2, 228, 229, 5, 46, 24, 7, 229, 231, 3, 2, 2, 2, 230, 196, 3,
	2, 2, 2, 230, 198, 3, 2, 2, 2, 230, 199, 3, 2, 2, 2, 230, 215, 3, 2, 2,
	2, 230, 227, 3, 2, 2, 2, 231, 252, 3, 2, 2, 2, 232, 233, 12, 6, 2, 2, 233,
	234, 5, 2, 2, 2, 234, 235, 5, 46, 24, 7, 235, 251, 3, 2, 2, 2, 236, 237,
	12, 5, 2, 2, 237, 238, 5, 4, 3, 2, 238, 239, 5, 46, 24, 6, 239, 251, 3,
	2, 2, 2, 240, 241, 12, 4, 2, 2, 241, 242, 5, 6, 4, 2, 242, 243, 5, 46,
	24, 5, 243, 251, 3, 2, 2, 2, 244, 245, 12, 3, 2, 2, 245, 246, 5, 8, 5,
	2, 246, 247, 5, 46, 24, 4, 247, 251, 3, 2, 2, 2, 248, 249, 12, 8, 2, 2,
	249, 251, 5, 10, 6, 2, 250, 232, 3, 2, 2, 2, 250, 236, 3, 2, 2, 2, 250,
	240, 3, 2, 2, 2, 250, 244, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 251, 254,
	3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 47, 3, 2,
	2, 2, 254, 252, 3, 2, 2, 2, 255, 275, 5, 14, 8, 2, 256, 258, 9, 7, 2, 2,
	257, 256, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259,
	261, 5, 58, 30, 2, 260, 262, 5, 30, 16, 2, 261, 260, 3, 2, 2, 2, 261, 262,
	3, 2, 2, 2, 262, 264, 3, 2, 2, 2, 263, 265, 5, 34, 18, 2, 264, 263, 3,
	2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 275, 3, 2, 2, 2, 266, 275, 7, 47, 2,
	2, 267, 275, 7, 20, 2, 2, 268, 269, 7, 3, 2, 2, 269, 270, 5, 46, 24, 2,
	270, 272, 7, 5, 2, 2, 271, 273, 5, 34, 18, 2, 272, 271, 3, 2, 2, 2, 272,
	273, 3, 2, 2, 2, 273, 275, 3, 2, 2, 2, 274, 255, 3, 2, 2, 2, 274, 257,
	3, 2, 2, 2, 274, 266, 3, 2, 2, 2, 274, 267, 3, 2, 2, 2, 274, 268, 3, 2,
	2, 2, 275, 49, 3, 2, 2, 2, 276, 277, 7, 14, 2, 2, 277, 285, 5, 52, 27,
	2, 278, 279, 7, 18, 2, 2, 279, 285, 5, 52, 27, 2, 280, 281, 7, 15, 2, 2,
	281, 285, 7, 42, 2, 2, 282, 285, 7, 16, 2, 2, 283, 285, 5, 52, 27, 2, 284,
	276, 3, 2, 2, 2, 284, 278, 3, 2, 2, 2, 284, 280, 3, 2, 2, 2, 284, 282,
	3, 2, 2, 2, 284, 283, 3, 2, 2, 2, 285, 51, 3, 2, 2, 2, 286, 290, 5, 54,
	28, 2, 287, 289, 5, 56, 29, 2, 288, 287, 3, 2, 2, 2, 289, 292, 3, 2, 2,
	2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292,
	290, 3, 2, 2, 2, 293, 295, 7, 43, 2, 2, 294, 293, 3, 2, 2, 2, 294, 295,
	3, 2, 2, 2, 295, 53, 3, 2, 2, 2, 296, 303, 7, 17, 2, 2, 297, 303, 7, 18,
	2, 2, 298, 303, 7, 42, 2, 2, 299, 303, 5, 32, 17, 2, 300, 303, 5, 34, 18,
	2, 301, 303, 7, 24, 2, 2, 302, 296, 3, 2, 2, 2, 302, 297, 3, 2, 2, 2, 302,
	298, 3, 2, 2, 2, 302, 299, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 301,
	3, 2, 2, 2, 303, 55, 3, 2, 2, 2, 304, 305, 7, 41, 2, 2, 305, 311, 7, 42,
	2, 2, 306, 307, 7, 41, 2, 2, 307, 311, 7, 45, 2, 2, 308, 311, 5, 32, 17,
	2, 309, 311, 5, 34, 18, 2, 310, 304, 3, 2, 2, 2, 310, 306, 3, 2, 2, 2,
	310, 308, 3, 2, 2, 2, 310, 309, 3, 2, 2, 2, 311, 57, 3, 2, 2, 2, 312, 316,
	5, 60, 31, 2, 313, 315, 5, 62, 32, 2, 314, 313, 3, 2, 2, 2, 315, 318, 3,
	2, 2, 2, 316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 59, 3, 2, 2,
	2, 318, 316, 3, 2, 2, 2, 319, 320, 9, 8, 2, 2, 320, 61, 3, 2, 2, 2, 321,
	322, 7, 41, 2, 2, 322, 328, 7, 42, 2, 2, 323, 324, 7, 41, 2, 2, 324, 328,
	7, 45, 2, 2, 325, 328, 7, 24, 2, 2, 326, 328, 5, 32, 17, 2, 327, 321, 3,
	2, 2, 2, 327, 323, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 326, 3, 2, 2,
	2, 328, 63, 3, 2, 2, 2, 329, 330, 7, 10, 2, 2, 330, 334, 5, 18, 10, 2,
	331, 332, 7, 10, 2, 2, 332, 334, 7, 42, 2, 2, 333, 329, 3, 2, 2, 2, 333,
	331, 3, 2, 2, 2, 334, 65, 3, 2, 2, 2, 41, 79, 84, 90, 92, 96, 101, 114,
	117, 121, 125, 128, 134, 137, 162, 168, 170, 177, 185, 201, 209, 212, 221,
	224, 230, 250, 252, 257, 261, 264, 272, 274, 284, 290, 294, 302, 310, 316,
	327, 333,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'{'", "'}'", "':'", "';'", "'post'", "", "'where'",
	"'else'", "'var'", "", "", "'$root'", "'root'", "'dest'", "", "", "'['",
	"']'", "", "'*'", "'/'", "'-'", "'+'", "'~='", "'='", "'>'", "'>='", "'<'",
	"'<='", "'~'", "'and'", "'or'", "'?'", "'def'", "'required'", "'.'", "",
	"'!'", "'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "IF", "WHERE", "ELSE", "VAR", "OBJ",
	"THIS", "ROOT_INPUT", "ROOT", "DEST", "BOOL", "NEWLINE", "LISTOPEN", "LISTCLOSE",
	"WILDCARD", "MUL", "DIV", "SUB", "ADD", "NEQ", "EQ", "GT", "GTEQ", "LT",
	"LTEQ", "NOT", "AND", "OR", "NOTNIL", "DEF", "REQUIRED", "DELIM", "TOKEN",
	"OWMOD", "CTX", "INTEGER", "COMMENT", "STRING", "WS", "UNKNOWN",
}

var ruleNames = []string{
	"bioperator1", "bioperator2", "bioperator3", "bioperator4", "postunoperator",
	"preunoperator", "floatingPoint", "root", "projectorDef", "argAlias", "conditionBlock",
	"ifBlock", "elseBlock", "inlineCondition", "inlineFilter", "index", "arrayMod",
	"block", "mapping", "comment", "condition", "filter", "expression", "source",
	"target", "targetPath", "targetPathHead", "targetPathSegment", "sourcePath",
	"sourcePathHead", "sourcePathSegment", "postProcess",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type WhistleParser struct {
	*antlr.BaseParser
}

func NewWhistleParser(input antlr.TokenStream) *WhistleParser {
	this := new(WhistleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Whistle.g4"

	return this
}

// WhistleParser tokens.
const (
	WhistleParserEOF        = antlr.TokenEOF
	WhistleParserT__0       = 1
	WhistleParserT__1       = 2
	WhistleParserT__2       = 3
	WhistleParserT__3       = 4
	WhistleParserT__4       = 5
	WhistleParserT__5       = 6
	WhistleParserT__6       = 7
	WhistleParserT__7       = 8
	WhistleParserIF         = 9
	WhistleParserWHERE      = 10
	WhistleParserELSE       = 11
	WhistleParserVAR        = 12
	WhistleParserOBJ        = 13
	WhistleParserTHIS       = 14
	WhistleParserROOT_INPUT = 15
	WhistleParserROOT       = 16
	WhistleParserDEST       = 17
	WhistleParserBOOL       = 18
	WhistleParserNEWLINE    = 19
	WhistleParserLISTOPEN   = 20
	WhistleParserLISTCLOSE  = 21
	WhistleParserWILDCARD   = 22
	WhistleParserMUL        = 23
	WhistleParserDIV        = 24
	WhistleParserSUB        = 25
	WhistleParserADD        = 26
	WhistleParserNEQ        = 27
	WhistleParserEQ         = 28
	WhistleParserGT         = 29
	WhistleParserGTEQ       = 30
	WhistleParserLT         = 31
	WhistleParserLTEQ       = 32
	WhistleParserNOT        = 33
	WhistleParserAND        = 34
	WhistleParserOR         = 35
	WhistleParserNOTNIL     = 36
	WhistleParserDEF        = 37
	WhistleParserREQUIRED   = 38
	WhistleParserDELIM      = 39
	WhistleParserTOKEN      = 40
	WhistleParserOWMOD      = 41
	WhistleParserCTX        = 42
	WhistleParserINTEGER    = 43
	WhistleParserCOMMENT    = 44
	WhistleParserSTRING     = 45
	WhistleParserWS         = 46
	WhistleParserUNKNOWN    = 47
)

// WhistleParser rules.
const (
	WhistleParserRULE_bioperator1       = 0
	WhistleParserRULE_bioperator2       = 1
	WhistleParserRULE_bioperator3       = 2
	WhistleParserRULE_bioperator4       = 3
	WhistleParserRULE_postunoperator    = 4
	WhistleParserRULE_preunoperator     = 5
	WhistleParserRULE_floatingPoint     = 6
	WhistleParserRULE_root              = 7
	WhistleParserRULE_projectorDef      = 8
	WhistleParserRULE_argAlias          = 9
	WhistleParserRULE_conditionBlock    = 10
	WhistleParserRULE_ifBlock           = 11
	WhistleParserRULE_elseBlock         = 12
	WhistleParserRULE_inlineCondition   = 13
	WhistleParserRULE_inlineFilter      = 14
	WhistleParserRULE_index             = 15
	WhistleParserRULE_arrayMod          = 16
	WhistleParserRULE_block             = 17
	WhistleParserRULE_mapping           = 18
	WhistleParserRULE_comment           = 19
	WhistleParserRULE_condition         = 20
	WhistleParserRULE_filter            = 21
	WhistleParserRULE_expression        = 22
	WhistleParserRULE_source            = 23
	WhistleParserRULE_target            = 24
	WhistleParserRULE_targetPath        = 25
	WhistleParserRULE_targetPathHead    = 26
	WhistleParserRULE_targetPathSegment = 27
	WhistleParserRULE_sourcePath        = 28
	WhistleParserRULE_sourcePathHead    = 29
	WhistleParserRULE_sourcePathSegment = 30
	WhistleParserRULE_postProcess       = 31
)

// IBioperator1Context is an interface to support dynamic dispatch.
type IBioperator1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBioperator1Context differentiates from other interfaces.
	IsBioperator1Context()
}

type Bioperator1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBioperator1Context() *Bioperator1Context {
	var p = new(Bioperator1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_bioperator1
	return p
}

func (*Bioperator1Context) IsBioperator1Context() {}

func NewBioperator1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bioperator1Context {
	var p = new(Bioperator1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_bioperator1

	return p
}

func (s *Bioperator1Context) GetParser() antlr.Parser { return s.parser }

func (s *Bioperator1Context) MUL() antlr.TerminalNode {
	return s.GetToken(WhistleParserMUL, 0)
}

func (s *Bioperator1Context) DIV() antlr.TerminalNode {
	return s.GetToken(WhistleParserDIV, 0)
}

func (s *Bioperator1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bioperator1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bioperator1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterBioperator1(s)
	}
}

func (s *Bioperator1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitBioperator1(s)
	}
}

func (s *Bioperator1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitBioperator1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Bioperator1() (localctx IBioperator1Context) {
	localctx = NewBioperator1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WhistleParserRULE_bioperator1)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WhistleParserMUL || _la == WhistleParserDIV) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBioperator2Context is an interface to support dynamic dispatch.
type IBioperator2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBioperator2Context differentiates from other interfaces.
	IsBioperator2Context()
}

type Bioperator2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBioperator2Context() *Bioperator2Context {
	var p = new(Bioperator2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_bioperator2
	return p
}

func (*Bioperator2Context) IsBioperator2Context() {}

func NewBioperator2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bioperator2Context {
	var p = new(Bioperator2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_bioperator2

	return p
}

func (s *Bioperator2Context) GetParser() antlr.Parser { return s.parser }

func (s *Bioperator2Context) ADD() antlr.TerminalNode {
	return s.GetToken(WhistleParserADD, 0)
}

func (s *Bioperator2Context) SUB() antlr.TerminalNode {
	return s.GetToken(WhistleParserSUB, 0)
}

func (s *Bioperator2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bioperator2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bioperator2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterBioperator2(s)
	}
}

func (s *Bioperator2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitBioperator2(s)
	}
}

func (s *Bioperator2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitBioperator2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Bioperator2() (localctx IBioperator2Context) {
	localctx = NewBioperator2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WhistleParserRULE_bioperator2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WhistleParserSUB || _la == WhistleParserADD) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBioperator3Context is an interface to support dynamic dispatch.
type IBioperator3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBioperator3Context differentiates from other interfaces.
	IsBioperator3Context()
}

type Bioperator3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBioperator3Context() *Bioperator3Context {
	var p = new(Bioperator3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_bioperator3
	return p
}

func (*Bioperator3Context) IsBioperator3Context() {}

func NewBioperator3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bioperator3Context {
	var p = new(Bioperator3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_bioperator3

	return p
}

func (s *Bioperator3Context) GetParser() antlr.Parser { return s.parser }

func (s *Bioperator3Context) EQ() antlr.TerminalNode {
	return s.GetToken(WhistleParserEQ, 0)
}

func (s *Bioperator3Context) NEQ() antlr.TerminalNode {
	return s.GetToken(WhistleParserNEQ, 0)
}

func (s *Bioperator3Context) GT() antlr.TerminalNode {
	return s.GetToken(WhistleParserGT, 0)
}

func (s *Bioperator3Context) GTEQ() antlr.TerminalNode {
	return s.GetToken(WhistleParserGTEQ, 0)
}

func (s *Bioperator3Context) LT() antlr.TerminalNode {
	return s.GetToken(WhistleParserLT, 0)
}

func (s *Bioperator3Context) LTEQ() antlr.TerminalNode {
	return s.GetToken(WhistleParserLTEQ, 0)
}

func (s *Bioperator3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bioperator3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bioperator3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterBioperator3(s)
	}
}

func (s *Bioperator3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitBioperator3(s)
	}
}

func (s *Bioperator3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitBioperator3(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Bioperator3() (localctx IBioperator3Context) {
	localctx = NewBioperator3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WhistleParserRULE_bioperator3)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(WhistleParserNEQ-27))|(1<<(WhistleParserEQ-27))|(1<<(WhistleParserGT-27))|(1<<(WhistleParserGTEQ-27))|(1<<(WhistleParserLT-27))|(1<<(WhistleParserLTEQ-27)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBioperator4Context is an interface to support dynamic dispatch.
type IBioperator4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBioperator4Context differentiates from other interfaces.
	IsBioperator4Context()
}

type Bioperator4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBioperator4Context() *Bioperator4Context {
	var p = new(Bioperator4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_bioperator4
	return p
}

func (*Bioperator4Context) IsBioperator4Context() {}

func NewBioperator4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bioperator4Context {
	var p = new(Bioperator4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_bioperator4

	return p
}

func (s *Bioperator4Context) GetParser() antlr.Parser { return s.parser }

func (s *Bioperator4Context) AND() antlr.TerminalNode {
	return s.GetToken(WhistleParserAND, 0)
}

func (s *Bioperator4Context) OR() antlr.TerminalNode {
	return s.GetToken(WhistleParserOR, 0)
}

func (s *Bioperator4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bioperator4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bioperator4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterBioperator4(s)
	}
}

func (s *Bioperator4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitBioperator4(s)
	}
}

func (s *Bioperator4Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitBioperator4(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Bioperator4() (localctx IBioperator4Context) {
	localctx = NewBioperator4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WhistleParserRULE_bioperator4)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WhistleParserAND || _la == WhistleParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPostunoperatorContext is an interface to support dynamic dispatch.
type IPostunoperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostunoperatorContext differentiates from other interfaces.
	IsPostunoperatorContext()
}

type PostunoperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostunoperatorContext() *PostunoperatorContext {
	var p = new(PostunoperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_postunoperator
	return p
}

func (*PostunoperatorContext) IsPostunoperatorContext() {}

func NewPostunoperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostunoperatorContext {
	var p = new(PostunoperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_postunoperator

	return p
}

func (s *PostunoperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *PostunoperatorContext) NOTNIL() antlr.TerminalNode {
	return s.GetToken(WhistleParserNOTNIL, 0)
}

func (s *PostunoperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostunoperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostunoperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterPostunoperator(s)
	}
}

func (s *PostunoperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitPostunoperator(s)
	}
}

func (s *PostunoperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitPostunoperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Postunoperator() (localctx IPostunoperatorContext) {
	localctx = NewPostunoperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WhistleParserRULE_postunoperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(WhistleParserNOTNIL)
	}

	return localctx
}

// IPreunoperatorContext is an interface to support dynamic dispatch.
type IPreunoperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreunoperatorContext differentiates from other interfaces.
	IsPreunoperatorContext()
}

type PreunoperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreunoperatorContext() *PreunoperatorContext {
	var p = new(PreunoperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_preunoperator
	return p
}

func (*PreunoperatorContext) IsPreunoperatorContext() {}

func NewPreunoperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreunoperatorContext {
	var p = new(PreunoperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_preunoperator

	return p
}

func (s *PreunoperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *PreunoperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(WhistleParserNOT, 0)
}

func (s *PreunoperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreunoperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreunoperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterPreunoperator(s)
	}
}

func (s *PreunoperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitPreunoperator(s)
	}
}

func (s *PreunoperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitPreunoperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Preunoperator() (localctx IPreunoperatorContext) {
	localctx = NewPreunoperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WhistleParserRULE_preunoperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(WhistleParserNOT)
	}

	return localctx
}

// IFloatingPointContext is an interface to support dynamic dispatch.
type IFloatingPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatingPointContext differentiates from other interfaces.
	IsFloatingPointContext()
}

type FloatingPointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatingPointContext() *FloatingPointContext {
	var p = new(FloatingPointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_floatingPoint
	return p
}

func (*FloatingPointContext) IsFloatingPointContext() {}

func NewFloatingPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatingPointContext {
	var p = new(FloatingPointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_floatingPoint

	return p
}

func (s *FloatingPointContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatingPointContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WhistleParserINTEGER)
}

func (s *FloatingPointContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WhistleParserINTEGER, i)
}

func (s *FloatingPointContext) SUB() antlr.TerminalNode {
	return s.GetToken(WhistleParserSUB, 0)
}

func (s *FloatingPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatingPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterFloatingPoint(s)
	}
}

func (s *FloatingPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitFloatingPoint(s)
	}
}

func (s *FloatingPointContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitFloatingPoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) FloatingPoint() (localctx IFloatingPointContext) {
	localctx = NewFloatingPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WhistleParserRULE_floatingPoint)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserSUB {
		{
			p.SetState(76)
			p.Match(WhistleParserSUB)
		}

	}
	{
		p.SetState(79)
		p.Match(WhistleParserINTEGER)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(80)
			p.Match(WhistleParserDELIM)
		}
		{
			p.SetState(81)
			p.Match(WhistleParserINTEGER)
		}

	}

	return localctx
}

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(WhistleParserEOF, 0)
}

func (s *RootContext) AllMapping() []IMappingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMappingContext)(nil)).Elem())
	var tst = make([]IMappingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMappingContext)
		}
	}

	return tst
}

func (s *RootContext) Mapping(i int) IMappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMappingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMappingContext)
}

func (s *RootContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *RootContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *RootContext) AllProjectorDef() []IProjectorDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProjectorDefContext)(nil)).Elem())
	var tst = make([]IProjectorDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProjectorDefContext)
		}
	}

	return tst
}

func (s *RootContext) ProjectorDef(i int) IProjectorDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectorDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProjectorDefContext)
}

func (s *RootContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(WhistleParserNEWLINE)
}

func (s *RootContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(WhistleParserNEWLINE, i)
}

func (s *RootContext) PostProcess() IPostProcessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostProcessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostProcessContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WhistleParserRULE_root)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(88)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case WhistleParserVAR, WhistleParserOBJ, WhistleParserTHIS, WhistleParserROOT_INPUT, WhistleParserROOT, WhistleParserLISTOPEN, WhistleParserWILDCARD, WhistleParserTOKEN:
				{
					p.SetState(84)
					p.Mapping()
				}

			case WhistleParserCOMMENT:
				{
					p.SetState(85)
					p.Comment()
				}

			case WhistleParserDEF:
				{
					p.SetState(86)
					p.ProjectorDef()
				}

			case WhistleParserNEWLINE:
				{
					p.SetState(87)
					p.Match(WhistleParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserT__7 {
		{
			p.SetState(93)
			p.PostProcess()
		}

	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WhistleParserNEWLINE {
		{
			p.SetState(96)
			p.Match(WhistleParserNEWLINE)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(102)
		p.Match(WhistleParserEOF)
	}

	return localctx
}

// IProjectorDefContext is an interface to support dynamic dispatch.
type IProjectorDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectorDefContext differentiates from other interfaces.
	IsProjectorDefContext()
}

type ProjectorDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectorDefContext() *ProjectorDefContext {
	var p = new(ProjectorDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_projectorDef
	return p
}

func (*ProjectorDefContext) IsProjectorDefContext() {}

func NewProjectorDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectorDefContext {
	var p = new(ProjectorDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_projectorDef

	return p
}

func (s *ProjectorDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectorDefContext) DEF() antlr.TerminalNode {
	return s.GetToken(WhistleParserDEF, 0)
}

func (s *ProjectorDefContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *ProjectorDefContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProjectorDefContext) AllArgAlias() []IArgAliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgAliasContext)(nil)).Elem())
	var tst = make([]IArgAliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgAliasContext)
		}
	}

	return tst
}

func (s *ProjectorDefContext) ArgAlias(i int) IArgAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgAliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgAliasContext)
}

func (s *ProjectorDefContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(WhistleParserNEWLINE)
}

func (s *ProjectorDefContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(WhistleParserNEWLINE, i)
}

func (s *ProjectorDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectorDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectorDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterProjectorDef(s)
	}
}

func (s *ProjectorDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitProjectorDef(s)
	}
}

func (s *ProjectorDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitProjectorDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) ProjectorDef() (localctx IProjectorDefContext) {
	localctx = NewProjectorDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, WhistleParserRULE_projectorDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(WhistleParserDEF)
	}
	{
		p.SetState(105)
		p.Match(WhistleParserTOKEN)
	}
	{
		p.SetState(106)
		p.Match(WhistleParserT__0)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserREQUIRED || _la == WhistleParserTOKEN {
		{
			p.SetState(107)
			p.ArgAlias()
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == WhistleParserT__1 {
			{
				p.SetState(108)
				p.Match(WhistleParserT__1)
			}
			{
				p.SetState(109)
				p.ArgAlias()
			}

			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(117)
		p.Match(WhistleParserT__2)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserNEWLINE {
		{
			p.SetState(118)
			p.Match(WhistleParserNEWLINE)
		}

	}
	{
		p.SetState(121)
		p.Block()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(122)
			p.Match(WhistleParserNEWLINE)
		}

	}

	return localctx
}

// IArgAliasContext is an interface to support dynamic dispatch.
type IArgAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgAliasContext differentiates from other interfaces.
	IsArgAliasContext()
}

type ArgAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgAliasContext() *ArgAliasContext {
	var p = new(ArgAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_argAlias
	return p
}

func (*ArgAliasContext) IsArgAliasContext() {}

func NewArgAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgAliasContext {
	var p = new(ArgAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_argAlias

	return p
}

func (s *ArgAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgAliasContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *ArgAliasContext) REQUIRED() antlr.TerminalNode {
	return s.GetToken(WhistleParserREQUIRED, 0)
}

func (s *ArgAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterArgAlias(s)
	}
}

func (s *ArgAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitArgAlias(s)
	}
}

func (s *ArgAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitArgAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) ArgAlias() (localctx IArgAliasContext) {
	localctx = NewArgAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, WhistleParserRULE_argAlias)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserREQUIRED {
		{
			p.SetState(125)
			p.Match(WhistleParserREQUIRED)
		}

	}
	{
		p.SetState(128)
		p.Match(WhistleParserTOKEN)
	}

	return localctx
}

// IConditionBlockContext is an interface to support dynamic dispatch.
type IConditionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionBlockContext differentiates from other interfaces.
	IsConditionBlockContext()
}

type ConditionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionBlockContext() *ConditionBlockContext {
	var p = new(ConditionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_conditionBlock
	return p
}

func (*ConditionBlockContext) IsConditionBlockContext() {}

func NewConditionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionBlockContext {
	var p = new(ConditionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_conditionBlock

	return p
}

func (s *ConditionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionBlockContext) IfBlock() IIfBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *ConditionBlockContext) ElseBlock() IElseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *ConditionBlockContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WhistleParserNEWLINE, 0)
}

func (s *ConditionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterConditionBlock(s)
	}
}

func (s *ConditionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitConditionBlock(s)
	}
}

func (s *ConditionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitConditionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) ConditionBlock() (localctx IConditionBlockContext) {
	localctx = NewConditionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, WhistleParserRULE_conditionBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.IfBlock()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserELSE {
		{
			p.SetState(131)
			p.ElseBlock()
		}

	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(134)
			p.Match(WhistleParserNEWLINE)
		}

	}

	return localctx
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_ifBlock
	return p
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfBlockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (s *IfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, WhistleParserRULE_ifBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Condition()
	}
	{
		p.SetState(138)
		p.Block()
	}

	return localctx
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_elseBlock
	return p
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(WhistleParserELSE, 0)
}

func (s *ElseBlockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (s *ElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, WhistleParserRULE_elseBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(WhistleParserELSE)
	}
	{
		p.SetState(141)
		p.Block()
	}

	return localctx
}

// IInlineConditionContext is an interface to support dynamic dispatch.
type IInlineConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineConditionContext differentiates from other interfaces.
	IsInlineConditionContext()
}

type InlineConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineConditionContext() *InlineConditionContext {
	var p = new(InlineConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_inlineCondition
	return p
}

func (*InlineConditionContext) IsInlineConditionContext() {}

func NewInlineConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineConditionContext {
	var p = new(InlineConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_inlineCondition

	return p
}

func (s *InlineConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineConditionContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *InlineConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterInlineCondition(s)
	}
}

func (s *InlineConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitInlineCondition(s)
	}
}

func (s *InlineConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitInlineCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) InlineCondition() (localctx IInlineConditionContext) {
	localctx = NewInlineConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, WhistleParserRULE_inlineCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(WhistleParserT__0)
	}
	{
		p.SetState(144)
		p.Condition()
	}
	{
		p.SetState(145)
		p.Match(WhistleParserT__2)
	}

	return localctx
}

// IInlineFilterContext is an interface to support dynamic dispatch.
type IInlineFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineFilterContext differentiates from other interfaces.
	IsInlineFilterContext()
}

type InlineFilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFilterContext() *InlineFilterContext {
	var p = new(InlineFilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_inlineFilter
	return p
}

func (*InlineFilterContext) IsInlineFilterContext() {}

func NewInlineFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFilterContext {
	var p = new(InlineFilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_inlineFilter

	return p
}

func (s *InlineFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFilterContext) LISTOPEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTOPEN, 0)
}

func (s *InlineFilterContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *InlineFilterContext) LISTCLOSE() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTCLOSE, 0)
}

func (s *InlineFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterInlineFilter(s)
	}
}

func (s *InlineFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitInlineFilter(s)
	}
}

func (s *InlineFilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitInlineFilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) InlineFilter() (localctx IInlineFilterContext) {
	localctx = NewInlineFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, WhistleParserRULE_inlineFilter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(WhistleParserLISTOPEN)
	}
	{
		p.SetState(148)
		p.Filter()
	}
	{
		p.SetState(149)
		p.Match(WhistleParserLISTCLOSE)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) LISTOPEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTOPEN, 0)
}

func (s *IndexContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WhistleParserINTEGER, 0)
}

func (s *IndexContext) LISTCLOSE() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTCLOSE, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, WhistleParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(WhistleParserLISTOPEN)
	}
	{
		p.SetState(152)
		p.Match(WhistleParserINTEGER)
	}
	{
		p.SetState(153)
		p.Match(WhistleParserLISTCLOSE)
	}

	return localctx
}

// IArrayModContext is an interface to support dynamic dispatch.
type IArrayModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayModContext differentiates from other interfaces.
	IsArrayModContext()
}

type ArrayModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayModContext() *ArrayModContext {
	var p = new(ArrayModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_arrayMod
	return p
}

func (*ArrayModContext) IsArrayModContext() {}

func NewArrayModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayModContext {
	var p = new(ArrayModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_arrayMod

	return p
}

func (s *ArrayModContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayModContext) LISTOPEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTOPEN, 0)
}

func (s *ArrayModContext) LISTCLOSE() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTCLOSE, 0)
}

func (s *ArrayModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterArrayMod(s)
	}
}

func (s *ArrayModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitArrayMod(s)
	}
}

func (s *ArrayModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitArrayMod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) ArrayMod() (localctx IArrayModContext) {
	localctx = NewArrayModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, WhistleParserRULE_arrayMod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(WhistleParserLISTOPEN)
	}
	{
		p.SetState(156)
		p.Match(WhistleParserLISTCLOSE)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(WhistleParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(WhistleParserNEWLINE, i)
}

func (s *BlockContext) AllMapping() []IMappingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMappingContext)(nil)).Elem())
	var tst = make([]IMappingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMappingContext)
		}
	}

	return tst
}

func (s *BlockContext) Mapping(i int) IMappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMappingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMappingContext)
}

func (s *BlockContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *BlockContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *BlockContext) AllConditionBlock() []IConditionBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionBlockContext)(nil)).Elem())
	var tst = make([]IConditionBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionBlockContext)
		}
	}

	return tst
}

func (s *BlockContext) ConditionBlock(i int) IConditionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionBlockContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, WhistleParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(WhistleParserT__3)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(159)
			p.Match(WhistleParserNEWLINE)
		}

	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WhistleParserIF)|(1<<WhistleParserVAR)|(1<<WhistleParserOBJ)|(1<<WhistleParserTHIS)|(1<<WhistleParserROOT_INPUT)|(1<<WhistleParserROOT)|(1<<WhistleParserNEWLINE)|(1<<WhistleParserLISTOPEN)|(1<<WhistleParserWILDCARD))) != 0) || _la == WhistleParserTOKEN || _la == WhistleParserCOMMENT {
		p.SetState(166)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case WhistleParserVAR, WhistleParserOBJ, WhistleParserTHIS, WhistleParserROOT_INPUT, WhistleParserROOT, WhistleParserLISTOPEN, WhistleParserWILDCARD, WhistleParserTOKEN:
			{
				p.SetState(162)
				p.Mapping()
			}

		case WhistleParserCOMMENT:
			{
				p.SetState(163)
				p.Comment()
			}

		case WhistleParserIF:
			{
				p.SetState(164)
				p.ConditionBlock()
			}

		case WhistleParserNEWLINE:
			{
				p.SetState(165)
				p.Match(WhistleParserNEWLINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
		p.Match(WhistleParserT__4)
	}

	return localctx
}

// IMappingContext is an interface to support dynamic dispatch.
type IMappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMappingContext differentiates from other interfaces.
	IsMappingContext()
}

type MappingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMappingContext() *MappingContext {
	var p = new(MappingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_mapping
	return p
}

func (*MappingContext) IsMappingContext() {}

func NewMappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MappingContext {
	var p = new(MappingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_mapping

	return p
}

func (s *MappingContext) GetParser() antlr.Parser { return s.parser }

func (s *MappingContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *MappingContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MappingContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *MappingContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WhistleParserNEWLINE, 0)
}

func (s *MappingContext) EOF() antlr.TerminalNode {
	return s.GetToken(WhistleParserEOF, 0)
}

func (s *MappingContext) InlineCondition() IInlineConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineConditionContext)
}

func (s *MappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterMapping(s)
	}
}

func (s *MappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitMapping(s)
	}
}

func (s *MappingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitMapping(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Mapping() (localctx IMappingContext) {
	localctx = NewMappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, WhistleParserRULE_mapping)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Target()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserT__0 {
		{
			p.SetState(174)
			p.InlineCondition()
		}

	}
	{
		p.SetState(177)
		p.Match(WhistleParserT__5)
	}
	{
		p.SetState(178)
		p.expression(0)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WhistleParserT__6:
		{
			p.SetState(179)
			p.Match(WhistleParserT__6)
		}

	case WhistleParserCOMMENT:
		{
			p.SetState(180)
			p.Comment()
		}

	case WhistleParserNEWLINE:
		{
			p.SetState(181)
			p.Match(WhistleParserNEWLINE)
		}

	case WhistleParserEOF:
		{
			p.SetState(182)
			p.Match(WhistleParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(WhistleParserCOMMENT, 0)
}

func (s *CommentContext) EOF() antlr.TerminalNode {
	return s.GetToken(WhistleParserEOF, 0)
}

func (s *CommentContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(WhistleParserNEWLINE, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, WhistleParserRULE_comment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(WhistleParserCOMMENT)
	}
	{
		p.SetState(186)
		_la = p.GetTokenStream().LA(1)

		if !(_la == WhistleParserEOF || _la == WhistleParserNEWLINE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) IF() antlr.TerminalNode {
	return s.GetToken(WhistleParserIF, 0)
}

func (s *ConditionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, WhistleParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(WhistleParserIF)
	}
	{
		p.SetState(189)
		p.expression(0)
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) WHERE() antlr.TerminalNode {
	return s.GetToken(WhistleParserWHERE, 0)
}

func (s *FilterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (s *FilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitFilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, WhistleParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(WhistleParserWHERE)
	}
	{
		p.SetState(192)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListInitializationContext struct {
	*ExpressionContext
}

func NewListInitializationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListInitializationContext {
	var p = new(ListInitializationContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ListInitializationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListInitializationContext) LISTOPEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTOPEN, 0)
}

func (s *ListInitializationContext) LISTCLOSE() antlr.TerminalNode {
	return s.GetToken(WhistleParserLISTCLOSE, 0)
}

func (s *ListInitializationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ListInitializationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListInitializationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterListInitialization(s)
	}
}

func (s *ListInitializationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitListInitialization(s)
	}
}

func (s *ListInitializationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitListInitialization(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprPreOpContext struct {
	*ExpressionContext
}

func NewExprPreOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprPreOpContext {
	var p = new(ExprPreOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprPreOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPreOpContext) Preunoperator() IPreunoperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreunoperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreunoperatorContext)
}

func (s *ExprPreOpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprPreOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterExprPreOp(s)
	}
}

func (s *ExprPreOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitExprPreOp(s)
	}
}

func (s *ExprPreOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitExprPreOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprPostOpContext struct {
	*ExpressionContext
}

func NewExprPostOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprPostOpContext {
	var p = new(ExprPostOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprPostOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPostOpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprPostOpContext) Postunoperator() IPostunoperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostunoperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostunoperatorContext)
}

func (s *ExprPostOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterExprPostOp(s)
	}
}

func (s *ExprPostOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitExprPostOp(s)
	}
}

func (s *ExprPostOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitExprPostOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAnonBlockContext struct {
	*ExpressionContext
}

func NewExprAnonBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAnonBlockContext {
	var p = new(ExprAnonBlockContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprAnonBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAnonBlockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ExprAnonBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterExprAnonBlock(s)
	}
}

func (s *ExprAnonBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitExprAnonBlock(s)
	}
}

func (s *ExprAnonBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitExprAnonBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprProjectionContext struct {
	*ExpressionContext
}

func NewExprProjectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprProjectionContext {
	var p = new(ExprProjectionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprProjectionContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *ExprProjectionContext) ArrayMod() IArrayModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayModContext)
}

func (s *ExprProjectionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprProjectionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterExprProjection(s)
	}
}

func (s *ExprProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitExprProjection(s)
	}
}

func (s *ExprProjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitExprProjection(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprBiOpContext struct {
	*ExpressionContext
}

func NewExprBiOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBiOpContext {
	var p = new(ExprBiOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprBiOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBiOpContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprBiOpContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprBiOpContext) Bioperator1() IBioperator1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBioperator1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBioperator1Context)
}

func (s *ExprBiOpContext) Bioperator2() IBioperator2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBioperator2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBioperator2Context)
}

func (s *ExprBiOpContext) Bioperator3() IBioperator3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBioperator3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBioperator3Context)
}

func (s *ExprBiOpContext) Bioperator4() IBioperator4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBioperator4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBioperator4Context)
}

func (s *ExprBiOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterExprBiOp(s)
	}
}

func (s *ExprBiOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitExprBiOp(s)
	}
}

func (s *ExprBiOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitExprBiOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprSourceContext struct {
	*ExpressionContext
}

func NewExprSourceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprSourceContext {
	var p = new(ExprSourceContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprSourceContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *ExprSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterExprSource(s)
	}
}

func (s *ExprSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitExprSource(s)
	}
}

func (s *ExprSourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitExprSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *WhistleParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, WhistleParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprSourceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(195)
			p.Source()
		}

	case 2:
		localctx = NewExprAnonBlockContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(196)
			p.Block()
		}

	case 3:
		localctx = NewExprProjectionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Match(WhistleParserTOKEN)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == WhistleParserLISTOPEN {
			{
				p.SetState(198)
				p.ArrayMod()
			}

		}
		{
			p.SetState(201)
			p.Match(WhistleParserT__0)
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WhistleParserT__0)|(1<<WhistleParserT__3)|(1<<WhistleParserVAR)|(1<<WhistleParserROOT_INPUT)|(1<<WhistleParserROOT)|(1<<WhistleParserDEST)|(1<<WhistleParserBOOL)|(1<<WhistleParserLISTOPEN)|(1<<WhistleParserSUB))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(WhistleParserNOT-33))|(1<<(WhistleParserTOKEN-33))|(1<<(WhistleParserINTEGER-33))|(1<<(WhistleParserSTRING-33)))) != 0) {
			{
				p.SetState(202)
				p.expression(0)
			}
			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == WhistleParserT__1 {
				{
					p.SetState(203)
					p.Match(WhistleParserT__1)
				}
				{
					p.SetState(204)
					p.expression(0)
				}

				p.SetState(209)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(212)
			p.Match(WhistleParserT__2)
		}

	case 4:
		localctx = NewListInitializationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(213)
			p.Match(WhistleParserLISTOPEN)
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WhistleParserT__0)|(1<<WhistleParserT__3)|(1<<WhistleParserVAR)|(1<<WhistleParserROOT_INPUT)|(1<<WhistleParserROOT)|(1<<WhistleParserDEST)|(1<<WhistleParserBOOL)|(1<<WhistleParserLISTOPEN)|(1<<WhistleParserSUB))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(WhistleParserNOT-33))|(1<<(WhistleParserTOKEN-33))|(1<<(WhistleParserINTEGER-33))|(1<<(WhistleParserSTRING-33)))) != 0) {
			{
				p.SetState(214)
				p.expression(0)
			}
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == WhistleParserT__1 {
				{
					p.SetState(215)
					p.Match(WhistleParserT__1)
				}
				{
					p.SetState(216)
					p.expression(0)
				}

				p.SetState(221)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(224)
			p.Match(WhistleParserLISTCLOSE)
		}

	case 5:
		localctx = NewExprPreOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(225)
			p.Preunoperator()
		}
		{
			p.SetState(226)
			p.expression(5)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprBiOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhistleParserRULE_expression)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(231)
					p.Bioperator1()
				}
				{
					p.SetState(232)
					p.expression(5)
				}

			case 2:
				localctx = NewExprBiOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhistleParserRULE_expression)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(235)
					p.Bioperator2()
				}
				{
					p.SetState(236)
					p.expression(4)
				}

			case 3:
				localctx = NewExprBiOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhistleParserRULE_expression)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(239)
					p.Bioperator3()
				}
				{
					p.SetState(240)
					p.expression(3)
				}

			case 4:
				localctx = NewExprBiOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhistleParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(243)
					p.Bioperator4()
				}
				{
					p.SetState(244)
					p.expression(2)
				}

			case 5:
				localctx = NewExprPostOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhistleParserRULE_expression)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(247)
					p.Postunoperator()
				}

			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) CopyFrom(ctx *SourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SourceConstStrContext struct {
	*SourceContext
}

func NewSourceConstStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceConstStrContext {
	var p = new(SourceConstStrContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SourceConstStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceConstStrContext) STRING() antlr.TerminalNode {
	return s.GetToken(WhistleParserSTRING, 0)
}

func (s *SourceConstStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourceConstStr(s)
	}
}

func (s *SourceConstStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourceConstStr(s)
	}
}

func (s *SourceConstStrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourceConstStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SourceProjectionContext struct {
	*SourceContext
}

func NewSourceProjectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceProjectionContext {
	var p = new(SourceProjectionContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SourceProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceProjectionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SourceProjectionContext) ArrayMod() IArrayModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayModContext)
}

func (s *SourceProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourceProjection(s)
	}
}

func (s *SourceProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourceProjection(s)
	}
}

func (s *SourceProjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourceProjection(s)

	default:
		return t.VisitChildren(s)
	}
}

type SourceConstNumContext struct {
	*SourceContext
}

func NewSourceConstNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceConstNumContext {
	var p = new(SourceConstNumContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SourceConstNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceConstNumContext) FloatingPoint() IFloatingPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatingPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatingPointContext)
}

func (s *SourceConstNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourceConstNum(s)
	}
}

func (s *SourceConstNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourceConstNum(s)
	}
}

func (s *SourceConstNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourceConstNum(s)

	default:
		return t.VisitChildren(s)
	}
}

type SourceInputContext struct {
	*SourceContext
}

func NewSourceInputContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceInputContext {
	var p = new(SourceInputContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SourceInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceInputContext) SourcePath() ISourcePathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourcePathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourcePathContext)
}

func (s *SourceInputContext) InlineFilter() IInlineFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineFilterContext)
}

func (s *SourceInputContext) ArrayMod() IArrayModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayModContext)
}

func (s *SourceInputContext) VAR() antlr.TerminalNode {
	return s.GetToken(WhistleParserVAR, 0)
}

func (s *SourceInputContext) DEST() antlr.TerminalNode {
	return s.GetToken(WhistleParserDEST, 0)
}

func (s *SourceInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourceInput(s)
	}
}

func (s *SourceInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourceInput(s)
	}
}

func (s *SourceInputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourceInput(s)

	default:
		return t.VisitChildren(s)
	}
}

type SourceConstBoolContext struct {
	*SourceContext
}

func NewSourceConstBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceConstBoolContext {
	var p = new(SourceConstBoolContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SourceConstBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceConstBoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(WhistleParserBOOL, 0)
}

func (s *SourceConstBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourceConstBool(s)
	}
}

func (s *SourceConstBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourceConstBool(s)
	}
}

func (s *SourceConstBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourceConstBool(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, WhistleParserRULE_source)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WhistleParserSUB, WhistleParserINTEGER:
		localctx = NewSourceConstNumContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.FloatingPoint()
		}

	case WhistleParserVAR, WhistleParserROOT_INPUT, WhistleParserROOT, WhistleParserDEST, WhistleParserTOKEN:
		localctx = NewSourceInputContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == WhistleParserVAR || _la == WhistleParserDEST {
			{
				p.SetState(254)
				_la = p.GetTokenStream().LA(1)

				if !(_la == WhistleParserVAR || _la == WhistleParserDEST) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(257)
			p.SourcePath()
		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(258)
				p.InlineFilter()
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(261)
				p.ArrayMod()
			}

		}

	case WhistleParserSTRING:
		localctx = NewSourceConstStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(WhistleParserSTRING)
		}

	case WhistleParserBOOL:
		localctx = NewSourceConstBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(265)
			p.Match(WhistleParserBOOL)
		}

	case WhistleParserT__0:
		localctx = NewSourceProjectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(266)
			p.Match(WhistleParserT__0)
		}
		{
			p.SetState(267)
			p.expression(0)
		}
		{
			p.SetState(268)
			p.Match(WhistleParserT__2)
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(269)
				p.ArrayMod()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) CopyFrom(ctx *TargetContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetVarContext struct {
	*TargetContext
}

func NewTargetVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetVarContext {
	var p = new(TargetVarContext)

	p.TargetContext = NewEmptyTargetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TargetContext))

	return p
}

func (s *TargetVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetVarContext) VAR() antlr.TerminalNode {
	return s.GetToken(WhistleParserVAR, 0)
}

func (s *TargetVarContext) TargetPath() ITargetPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetPathContext)
}

func (s *TargetVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetVar(s)
	}
}

func (s *TargetVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetVar(s)
	}
}

func (s *TargetVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetObjContext struct {
	*TargetContext
}

func NewTargetObjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetObjContext {
	var p = new(TargetObjContext)

	p.TargetContext = NewEmptyTargetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TargetContext))

	return p
}

func (s *TargetObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetObjContext) OBJ() antlr.TerminalNode {
	return s.GetToken(WhistleParserOBJ, 0)
}

func (s *TargetObjContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *TargetObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetObj(s)
	}
}

func (s *TargetObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetObj(s)
	}
}

func (s *TargetObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetObj(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetRootFieldContext struct {
	*TargetContext
}

func NewTargetRootFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetRootFieldContext {
	var p = new(TargetRootFieldContext)

	p.TargetContext = NewEmptyTargetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TargetContext))

	return p
}

func (s *TargetRootFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRootFieldContext) ROOT() antlr.TerminalNode {
	return s.GetToken(WhistleParserROOT, 0)
}

func (s *TargetRootFieldContext) TargetPath() ITargetPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetPathContext)
}

func (s *TargetRootFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetRootField(s)
	}
}

func (s *TargetRootFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetRootField(s)
	}
}

func (s *TargetRootFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetRootField(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetThisContext struct {
	*TargetContext
}

func NewTargetThisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetThisContext {
	var p = new(TargetThisContext)

	p.TargetContext = NewEmptyTargetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TargetContext))

	return p
}

func (s *TargetThisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetThisContext) THIS() antlr.TerminalNode {
	return s.GetToken(WhistleParserTHIS, 0)
}

func (s *TargetThisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetThis(s)
	}
}

func (s *TargetThisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetThis(s)
	}
}

func (s *TargetThisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetThis(s)

	default:
		return t.VisitChildren(s)
	}
}

type TargetFieldContext struct {
	*TargetContext
}

func NewTargetFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetFieldContext {
	var p = new(TargetFieldContext)

	p.TargetContext = NewEmptyTargetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TargetContext))

	return p
}

func (s *TargetFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetFieldContext) TargetPath() ITargetPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetPathContext)
}

func (s *TargetFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetField(s)
	}
}

func (s *TargetFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetField(s)
	}
}

func (s *TargetFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, WhistleParserRULE_target)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTargetVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Match(WhistleParserVAR)
		}
		{
			p.SetState(275)
			p.TargetPath()
		}

	case 2:
		localctx = NewTargetRootFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.Match(WhistleParserROOT)
		}
		{
			p.SetState(277)
			p.TargetPath()
		}

	case 3:
		localctx = NewTargetObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(278)
			p.Match(WhistleParserOBJ)
		}
		{
			p.SetState(279)
			p.Match(WhistleParserTOKEN)
		}

	case 4:
		localctx = NewTargetThisContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			p.Match(WhistleParserTHIS)
		}

	case 5:
		localctx = NewTargetFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(281)
			p.TargetPath()
		}

	}

	return localctx
}

// ITargetPathContext is an interface to support dynamic dispatch.
type ITargetPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetPathContext differentiates from other interfaces.
	IsTargetPathContext()
}

type TargetPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetPathContext() *TargetPathContext {
	var p = new(TargetPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_targetPath
	return p
}

func (*TargetPathContext) IsTargetPathContext() {}

func NewTargetPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetPathContext {
	var p = new(TargetPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_targetPath

	return p
}

func (s *TargetPathContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetPathContext) TargetPathHead() ITargetPathHeadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetPathHeadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetPathHeadContext)
}

func (s *TargetPathContext) AllTargetPathSegment() []ITargetPathSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITargetPathSegmentContext)(nil)).Elem())
	var tst = make([]ITargetPathSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITargetPathSegmentContext)
		}
	}

	return tst
}

func (s *TargetPathContext) TargetPathSegment(i int) ITargetPathSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetPathSegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITargetPathSegmentContext)
}

func (s *TargetPathContext) OWMOD() antlr.TerminalNode {
	return s.GetToken(WhistleParserOWMOD, 0)
}

func (s *TargetPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetPath(s)
	}
}

func (s *TargetPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetPath(s)
	}
}

func (s *TargetPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) TargetPath() (localctx ITargetPathContext) {
	localctx = NewTargetPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, WhistleParserRULE_targetPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.TargetPathHead()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WhistleParserLISTOPEN || _la == WhistleParserDELIM {
		{
			p.SetState(285)
			p.TargetPathSegment()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WhistleParserOWMOD {
		{
			p.SetState(291)
			p.Match(WhistleParserOWMOD)
		}

	}

	return localctx
}

// ITargetPathHeadContext is an interface to support dynamic dispatch.
type ITargetPathHeadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetPathHeadContext differentiates from other interfaces.
	IsTargetPathHeadContext()
}

type TargetPathHeadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetPathHeadContext() *TargetPathHeadContext {
	var p = new(TargetPathHeadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_targetPathHead
	return p
}

func (*TargetPathHeadContext) IsTargetPathHeadContext() {}

func NewTargetPathHeadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetPathHeadContext {
	var p = new(TargetPathHeadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_targetPathHead

	return p
}

func (s *TargetPathHeadContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetPathHeadContext) ROOT_INPUT() antlr.TerminalNode {
	return s.GetToken(WhistleParserROOT_INPUT, 0)
}

func (s *TargetPathHeadContext) ROOT() antlr.TerminalNode {
	return s.GetToken(WhistleParserROOT, 0)
}

func (s *TargetPathHeadContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *TargetPathHeadContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *TargetPathHeadContext) ArrayMod() IArrayModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayModContext)
}

func (s *TargetPathHeadContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(WhistleParserWILDCARD, 0)
}

func (s *TargetPathHeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetPathHeadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetPathHeadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetPathHead(s)
	}
}

func (s *TargetPathHeadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetPathHead(s)
	}
}

func (s *TargetPathHeadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetPathHead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) TargetPathHead() (localctx ITargetPathHeadContext) {
	localctx = NewTargetPathHeadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, WhistleParserRULE_targetPathHead)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(WhistleParserROOT_INPUT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(WhistleParserROOT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Match(WhistleParserTOKEN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Index()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(298)
			p.ArrayMod()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(299)
			p.Match(WhistleParserWILDCARD)
		}

	}

	return localctx
}

// ITargetPathSegmentContext is an interface to support dynamic dispatch.
type ITargetPathSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetPathSegmentContext differentiates from other interfaces.
	IsTargetPathSegmentContext()
}

type TargetPathSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetPathSegmentContext() *TargetPathSegmentContext {
	var p = new(TargetPathSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_targetPathSegment
	return p
}

func (*TargetPathSegmentContext) IsTargetPathSegmentContext() {}

func NewTargetPathSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetPathSegmentContext {
	var p = new(TargetPathSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_targetPathSegment

	return p
}

func (s *TargetPathSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetPathSegmentContext) DELIM() antlr.TerminalNode {
	return s.GetToken(WhistleParserDELIM, 0)
}

func (s *TargetPathSegmentContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *TargetPathSegmentContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WhistleParserINTEGER, 0)
}

func (s *TargetPathSegmentContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *TargetPathSegmentContext) ArrayMod() IArrayModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayModContext)
}

func (s *TargetPathSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetPathSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetPathSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterTargetPathSegment(s)
	}
}

func (s *TargetPathSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitTargetPathSegment(s)
	}
}

func (s *TargetPathSegmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitTargetPathSegment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) TargetPathSegment() (localctx ITargetPathSegmentContext) {
	localctx = NewTargetPathSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, WhistleParserRULE_targetPathSegment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Match(WhistleParserDELIM)
		}
		{
			p.SetState(303)
			p.Match(WhistleParserTOKEN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Match(WhistleParserDELIM)
		}
		{
			p.SetState(305)
			p.Match(WhistleParserINTEGER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(306)
			p.Index()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(307)
			p.ArrayMod()
		}

	}

	return localctx
}

// ISourcePathContext is an interface to support dynamic dispatch.
type ISourcePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourcePathContext differentiates from other interfaces.
	IsSourcePathContext()
}

type SourcePathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourcePathContext() *SourcePathContext {
	var p = new(SourcePathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_sourcePath
	return p
}

func (*SourcePathContext) IsSourcePathContext() {}

func NewSourcePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourcePathContext {
	var p = new(SourcePathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_sourcePath

	return p
}

func (s *SourcePathContext) GetParser() antlr.Parser { return s.parser }

func (s *SourcePathContext) SourcePathHead() ISourcePathHeadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourcePathHeadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourcePathHeadContext)
}

func (s *SourcePathContext) AllSourcePathSegment() []ISourcePathSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourcePathSegmentContext)(nil)).Elem())
	var tst = make([]ISourcePathSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourcePathSegmentContext)
		}
	}

	return tst
}

func (s *SourcePathContext) SourcePathSegment(i int) ISourcePathSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourcePathSegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourcePathSegmentContext)
}

func (s *SourcePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourcePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourcePathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourcePath(s)
	}
}

func (s *SourcePathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourcePath(s)
	}
}

func (s *SourcePathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourcePath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) SourcePath() (localctx ISourcePathContext) {
	localctx = NewSourcePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, WhistleParserRULE_sourcePath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.SourcePathHead()
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(311)
				p.SourcePathSegment()
			}

		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
	}

	return localctx
}

// ISourcePathHeadContext is an interface to support dynamic dispatch.
type ISourcePathHeadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourcePathHeadContext differentiates from other interfaces.
	IsSourcePathHeadContext()
}

type SourcePathHeadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourcePathHeadContext() *SourcePathHeadContext {
	var p = new(SourcePathHeadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_sourcePathHead
	return p
}

func (*SourcePathHeadContext) IsSourcePathHeadContext() {}

func NewSourcePathHeadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourcePathHeadContext {
	var p = new(SourcePathHeadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_sourcePathHead

	return p
}

func (s *SourcePathHeadContext) GetParser() antlr.Parser { return s.parser }

func (s *SourcePathHeadContext) ROOT_INPUT() antlr.TerminalNode {
	return s.GetToken(WhistleParserROOT_INPUT, 0)
}

func (s *SourcePathHeadContext) ROOT() antlr.TerminalNode {
	return s.GetToken(WhistleParserROOT, 0)
}

func (s *SourcePathHeadContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *SourcePathHeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourcePathHeadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourcePathHeadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourcePathHead(s)
	}
}

func (s *SourcePathHeadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourcePathHead(s)
	}
}

func (s *SourcePathHeadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourcePathHead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) SourcePathHead() (localctx ISourcePathHeadContext) {
	localctx = NewSourcePathHeadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, WhistleParserRULE_sourcePathHead)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-15)&-(0x1f+1)) == 0 && ((1<<uint((_la-15)))&((1<<(WhistleParserROOT_INPUT-15))|(1<<(WhistleParserROOT-15))|(1<<(WhistleParserTOKEN-15)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISourcePathSegmentContext is an interface to support dynamic dispatch.
type ISourcePathSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourcePathSegmentContext differentiates from other interfaces.
	IsSourcePathSegmentContext()
}

type SourcePathSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourcePathSegmentContext() *SourcePathSegmentContext {
	var p = new(SourcePathSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_sourcePathSegment
	return p
}

func (*SourcePathSegmentContext) IsSourcePathSegmentContext() {}

func NewSourcePathSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourcePathSegmentContext {
	var p = new(SourcePathSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_sourcePathSegment

	return p
}

func (s *SourcePathSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SourcePathSegmentContext) DELIM() antlr.TerminalNode {
	return s.GetToken(WhistleParserDELIM, 0)
}

func (s *SourcePathSegmentContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *SourcePathSegmentContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(WhistleParserINTEGER, 0)
}

func (s *SourcePathSegmentContext) WILDCARD() antlr.TerminalNode {
	return s.GetToken(WhistleParserWILDCARD, 0)
}

func (s *SourcePathSegmentContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *SourcePathSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourcePathSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourcePathSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterSourcePathSegment(s)
	}
}

func (s *SourcePathSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitSourcePathSegment(s)
	}
}

func (s *SourcePathSegmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitSourcePathSegment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) SourcePathSegment() (localctx ISourcePathSegmentContext) {
	localctx = NewSourcePathSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, WhistleParserRULE_sourcePathSegment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(WhistleParserDELIM)
		}
		{
			p.SetState(320)
			p.Match(WhistleParserTOKEN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.Match(WhistleParserDELIM)
		}
		{
			p.SetState(322)
			p.Match(WhistleParserINTEGER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(323)
			p.Match(WhistleParserWILDCARD)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(324)
			p.Index()
		}

	}

	return localctx
}

// IPostProcessContext is an interface to support dynamic dispatch.
type IPostProcessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostProcessContext differentiates from other interfaces.
	IsPostProcessContext()
}

type PostProcessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostProcessContext() *PostProcessContext {
	var p = new(PostProcessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhistleParserRULE_postProcess
	return p
}

func (*PostProcessContext) IsPostProcessContext() {}

func NewPostProcessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostProcessContext {
	var p = new(PostProcessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhistleParserRULE_postProcess

	return p
}

func (s *PostProcessContext) GetParser() antlr.Parser { return s.parser }

func (s *PostProcessContext) CopyFrom(ctx *PostProcessContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PostProcessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostProcessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PostProcessNameContext struct {
	*PostProcessContext
}

func NewPostProcessNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostProcessNameContext {
	var p = new(PostProcessNameContext)

	p.PostProcessContext = NewEmptyPostProcessContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PostProcessContext))

	return p
}

func (s *PostProcessNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostProcessNameContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(WhistleParserTOKEN, 0)
}

func (s *PostProcessNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterPostProcessName(s)
	}
}

func (s *PostProcessNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitPostProcessName(s)
	}
}

func (s *PostProcessNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitPostProcessName(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostProcessInlineContext struct {
	*PostProcessContext
}

func NewPostProcessInlineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostProcessInlineContext {
	var p = new(PostProcessInlineContext)

	p.PostProcessContext = NewEmptyPostProcessContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PostProcessContext))

	return p
}

func (s *PostProcessInlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostProcessInlineContext) ProjectorDef() IProjectorDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectorDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectorDefContext)
}

func (s *PostProcessInlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.EnterPostProcessInline(s)
	}
}

func (s *PostProcessInlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhistleListener); ok {
		listenerT.ExitPostProcessInline(s)
	}
}

func (s *PostProcessInlineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case WhistleVisitor:
		return t.VisitPostProcessInline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *WhistleParser) PostProcess() (localctx IPostProcessContext) {
	localctx = NewPostProcessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, WhistleParserRULE_postProcess)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPostProcessInlineContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(WhistleParserT__7)
		}
		{
			p.SetState(328)
			p.ProjectorDef()
		}

	case 2:
		localctx = NewPostProcessNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(WhistleParserT__7)
		}
		{
			p.SetState(330)
			p.Match(WhistleParserTOKEN)
		}

	}

	return localctx
}

func (p *WhistleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *WhistleParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
