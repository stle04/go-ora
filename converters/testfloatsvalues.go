package converters

/* This file is generated.
   go run .\generatefloat\main.go converters
*/

var TestFloatValue = []struct {
	SelectText string
	OracleText string
	Float      float64
	Integer    int64
	IsInteger  bool
	Binary     []byte
}{
	// { "0", "0", 0,  0, true,  []byte{ 128 } },  // 0.000000e+00
	// { "1", "1", 1,  1, true,  []byte{ 193,2 } },  // 1.000000e+00
	// { "10", "10", 10,  10, true,  []byte{ 193,11 } },  // 1.000000e+01
	// { "100", "100", 100,  100, true,  []byte{ 194,2 } },  // 1.000000e+02
	// { "1000", "1000", 1000,  1000, true,  []byte{ 194,11 } },  // 1.000000e+03
	{"10000000", "10000000", 1e+07, 10000000, true, []byte{196, 11}},                                               // 1.000000e+07
	{"1E+30", "1000000000000000000000000000000", 1e+30, 0, false, []byte{208, 2}},                                  // 1.000000e+30
	{"1E+125", "1.000000000000000000000000000000000E+125", 1e+125, 0, false, []byte{255, 11}},                      // 1.000000e+125
	{"0.1", ".1", 0.1, 0, false, []byte{192, 11}},                                                                  // 1.000000e-01
	{"0.01", ".01", 0.01, 0, false, []byte{192, 2}},                                                                // 1.000000e-02
	{"0.001", ".001", 0.001, 0, false, []byte{191, 11}},                                                            // 1.000000e-03
	{"0.0001", ".0001", 0.0001, 0, false, []byte{191, 2}},                                                          // 1.000000e-04
	{"0.00001", ".00001", 1e-05, 0, false, []byte{190, 11}},                                                        // 1.000000e-05
	{"0.000001", ".000001", 1e-06, 0, false, []byte{190, 2}},                                                       // 1.000000e-06
	{"1E+125", "1.000000000000000000000000000000000E+125", 1e+125, 0, false, []byte{255, 11}},                      // 1.000000e+125
	{"1E-125", "1.000000000000000000000000000000000E-125", 1e-125, 0, false, []byte{130, 11}},                      // 1.000000e-125
	{"-1E+125", "-1.00000000000000000000000000000000E+125", -1e+125, 0, false, []byte{0, 91, 102}},                 // -1.000000e+125
	{"-1E-125", "-1.00000000000000000000000000000000E-125", -1e-125, 0, false, []byte{125, 91, 102}},               // -1.000000e-125
	{"1.23456789e15", "1234567890000000", 1.23456789e+15, 1234567890000000, true, []byte{200, 13, 35, 57, 79, 91}}, // 1.234568e+15
	{"1.23456789e-15", ".00000000000000123456789", 1.23456789e-15, 0, false, []byte{185, 13, 35, 57, 79, 91}},      // 1.234568e-15
	{"1.234", "1.234", 1.234, 0, false, []byte{193, 2, 24, 41}},                                                    // 1.234000e+00
	{"12.34", "12.34", 12.34, 0, false, []byte{193, 13, 35}},                                                       // 1.234000e+01
	{"123.4", "123.4", 123.4, 0, false, []byte{194, 2, 24, 41}},                                                    // 1.234000e+02
	{"1234", "1234", 1234, 1234, true, []byte{194, 13, 35}},                                                        // 1.234000e+03
	{"12340", "12340", 12340, 12340, true, []byte{195, 2, 24, 41}},                                                 // 1.234000e+04
	{"123400", "123400", 123400, 123400, true, []byte{195, 13, 35}},                                                // 1.234000e+05
	{"1234000", "1234000", 1.234e+06, 1234000, true, []byte{196, 2, 24, 41}},                                       // 1.234000e+06
	{"12340000", "12340000", 1.234e+07, 12340000, true, []byte{196, 13, 35}},                                       // 1.234000e+07
	{"0.1234", ".1234", 0.1234, 0, false, []byte{192, 13, 35}},                                                     // 1.234000e-01
	{"0.01234", ".01234", 0.01234, 0, false, []byte{192, 2, 24, 41}},                                               // 1.234000e-02
	{"0.001234", ".001234", 0.001234, 0, false, []byte{191, 13, 35}},                                               // 1.234000e-03
	{"0.0001234", ".0001234", 0.0001234, 0, false, []byte{191, 2, 24, 41}},                                         // 1.234000e-04
	{"0.00001234", ".00001234", 1.234e-05, 0, false, []byte{190, 13, 35}},                                          // 1.234000e-05
	{"0.000001234", ".000001234", 1.234e-06, 0, false, []byte{190, 2, 24, 41}},                                     // 1.234000e-06
	{"-1.234", "-1.234", -1.234, 0, false, []byte{62, 100, 78, 61, 102}},                                           // -1.234000e+00
	{"-12.34", "-12.34", -12.34, 0, false, []byte{62, 89, 67, 102}},                                                // -1.234000e+01
	{"-123.4", "-123.4", -123.4, 0, false, []byte{61, 100, 78, 61, 102}},                                           // -1.234000e+02
	{"-1234", "-1234", -1234, -1234, true, []byte{61, 89, 67, 102}},                                                // -1.234000e+03
	{"-12340", "-12340", -12340, -12340, true, []byte{60, 100, 78, 61, 102}},                                       // -1.234000e+04
	{"-123400", "-123400", -123400, -123400, true, []byte{60, 89, 67, 102}},                                        // -1.234000e+05
	{"-1234000", "-1234000", -1.234e+06, -1234000, true, []byte{59, 100, 78, 61, 102}},                             // -1.234000e+06
	{"-12340000", "-12340000", -1.234e+07, -12340000, true, []byte{59, 89, 67, 102}},                               // -1.234000e+07
	{"-0.1234", "-.1234", -0.1234, 0, false, []byte{63, 89, 67, 102}},                                              // -1.234000e-01
	{"-1.234", "-1.234", -1.234, 0, false, []byte{62, 100, 78, 61, 102}},                                           // -1.234000e+00
	{"-12.34", "-12.34", -12.34, 0, false, []byte{62, 89, 67, 102}},                                                // -1.234000e+01
	{"-123.4", "-123.4", -123.4, 0, false, []byte{61, 100, 78, 61, 102}},                                           // -1.234000e+02
	{"-1234", "-1234", -1234, -1234, true, []byte{61, 89, 67, 102}},                                                // -1.234000e+03
	{"-12340", "-12340", -12340, -12340, true, []byte{60, 100, 78, 61, 102}},                                       // -1.234000e+04
	{"-123400", "-123400", -123400, -123400, true, []byte{60, 89, 67, 102}},                                        // -1.234000e+05
	{"-1234000", "-1234000", -1.234e+06, -1234000, true, []byte{59, 100, 78, 61, 102}},                             // -1.234000e+06
	{"9.8765", "9.8765", 9.8765, 0, false, []byte{193, 10, 88, 66}},                                                // 9.876500e+00
	{"98.765", "98.765", 98.765, 0, false, []byte{193, 99, 77, 51}},                                                // 9.876500e+01
	{"987.65", "987.65", 987.65, 0, false, []byte{194, 10, 88, 66}},                                                // 9.876500e+02
	{"9876.5", "9876.5", 9876.5, 0, false, []byte{194, 99, 77, 51}},                                                // 9.876500e+03
	{"98765", "98765", 98765, 98765, true, []byte{195, 10, 88, 66}},                                                // 9.876500e+04
	{"987650", "987650", 987650, 987650, true, []byte{195, 99, 77, 51}},                                            // 9.876500e+05
	{"9876500", "9876500", 9.8765e+06, 9876500, true, []byte{196, 10, 88, 66}},                                     // 9.876500e+06
	{"0.98765", ".98765", 0.98765, 0, false, []byte{192, 99, 77, 51}},                                              // 9.876500e-01
	{"0.098765", ".098765", 0.098765, 0, false, []byte{192, 10, 88, 66}},                                           // 9.876500e-02
	{"0.0098765", ".0098765", 0.0098765, 0, false, []byte{191, 99, 77, 51}},                                        // 9.876500e-03
	{"0.00098765", ".00098765", 0.00098765, 0, false, []byte{191, 10, 88, 66}},                                     // 9.876500e-04
	{"0.000098765", ".000098765", 9.8765e-05, 0, false, []byte{190, 99, 77, 51}},                                   // 9.876500e-05
	{"0.0000098765", ".0000098765", 9.8765e-06, 0, false, []byte{190, 10, 88, 66}},                                 // 9.876500e-06
	{"0.00000098765", ".00000098765", 9.8765e-07, 0, false, []byte{189, 99, 77, 51}},                               // 9.876500e-07
	{"-9.8765", "-9.8765", -9.8765, 0, false, []byte{62, 92, 14, 36, 102}},                                         // -9.876500e+00
	{"-98.765", "-98.765", -98.765, 0, false, []byte{62, 3, 25, 51, 102}},                                          // -9.876500e+01
	{"-987.65", "-987.65", -987.65, 0, false, []byte{61, 92, 14, 36, 102}},                                         // -9.876500e+02
	{"-9876.5", "-9876.5", -9876.5, 0, false, []byte{61, 3, 25, 51, 102}},                                          // -9.876500e+03
	{"-98765", "-98765", -98765, -98765, true, []byte{60, 92, 14, 36, 102}},                                        // -9.876500e+04
	{"-987650", "-987650", -987650, -987650, true, []byte{60, 3, 25, 51, 102}},                                     // -9.876500e+05
	{"-9876500", "-9876500", -9.8765e+06, -9876500, true, []byte{59, 92, 14, 36, 102}},                             // -9.876500e+06
	{"-98765000", "-98765000", -9.8765e+07, -98765000, true, []byte{59, 3, 25, 51, 102}},                           // -9.876500e+07
	{"-0.98765", "-.98765", -0.98765, 0, false, []byte{63, 3, 25, 51, 102}},                                        // -9.876500e-01
	{"-0.098765", "-.098765", -0.098765, 0, false, []byte{63, 92, 14, 36, 102}},                                    // -9.876500e-02
	{"-0.0098765", "-.0098765", -0.0098765, 0, false, []byte{64, 3, 25, 51, 102}},                                  // -9.876500e-03
	{"-0.00098765", "-.00098765", -0.00098765, 0, false, []byte{64, 92, 14, 36, 102}},                              // -9.876500e-04
	{"-0.000098765", "-.000098765", -9.8765e-05, 0, false, []byte{65, 3, 25, 51, 102}},                             // -9.876500e-05
	{"-0.0000098765", "-.0000098765", -9.8765e-06, 0, false, []byte{65, 92, 14, 36, 102}},                          // -9.876500e-06
	{"-0.00000098765", "-.00000098765", -9.8765e-07, 0, false, []byte{66, 3, 25, 51, 102}},                         // -9.876500e-07
	{"2*asin(1)", "3.1415926535897932384626433832795028842", 3.141592653589793, 0, false, []byte{193, 4, 15, 16, 93, 66, 36, 90, 80, 33, 39, 47, 27, 44, 39, 33, 80, 51, 29, 85, 21}}, // 3.141593e+00
	{"1/3", ".333333333333333333333333333333333333333", 0.3333333333333333, 0, false, []byte{192, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34}},    // 3.333333e-01
	{"-1/3", "-.33333333333333333333333333333333333333", -0.3333333333333333, 0, false, []byte{63, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68, 68}},   // -3.333333e-01
	{"9223372036854775807", "9223372036854775807", 9.223372036854776e+18, 9223372036854775807, true, []byte{202, 10, 23, 34, 73, 4, 69, 55, 78, 59, 8}},                               // 9.223372e+18
	{"-9223372036854775808", "-9223372036854775808", -9.223372036854776e+18, -9223372036854775808, true, []byte{53, 92, 79, 68, 29, 98, 33, 47, 24, 43, 93, 102}},                     // -9.223372e+18
}
