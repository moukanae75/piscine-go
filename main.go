package piscine

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// التحقق من وجود معاملات (الكود المصدري)
// 	if len(os.Args) < 2 {
// 		return // إذا لم يتم إدخال كود، لا تفعل شيئاً
// 	}

// 	// الحصول على الكود المصدري من المعاملات
// 	source := os.Args[1]

// 	// تنفيذ الكود
// 	interpretBrainfuck(source)
// }

// func interpretBrainfuck(source string) {
// 	// إنشاء مصفوفة البايتات (الذاكرة)
// 	memory := make([]byte, 2048)
	
// 	// المؤشر (الموقع الحالي في الذاكرة)
// 	pointer := 0
	
// 	// مؤشر الحلقة (لمعالجة الأقواس المربعة)
// 	i := 0
	
// 	// طول الكود المصدري
// 	length := len(source)
	
// 	// تنفيذ الكود حرفاً بحرفاً
// 	for i < length {
// 		char := source[i]
// 		println(char)
		
// 		switch char {
// 		case '>':
// 			// زيادة المؤشر (الانتقال إلى اليمين)
// 			pointer++
// 			// التأكد من عدم تجاوز حدود المصفوفة
// 			if pointer >= 2048 {
// 				pointer = 0 // العودة إلى البداية إذا وصلنا إلى النهاية
// 			}
// 		case '<':
// 			// إنقاص المؤشر (الانتقال إلى اليسار)
// 			pointer--
// 			// التأكد من عدم تجاوز حدود المصفوفة
// 			if pointer < 0 {
// 				pointer = 2047 // الانتقال إلى النهاية إذا كنا في البداية
// 			}
// 		case '+':
// 			// زيادة قيمة البايت الحالي
// 			memory[pointer]++
// 		case '-':
// 			// إنقاص قيمة البايت الحالي
// 			memory[pointer]--
// 		case '.':
// 			// طباعة الحرف المقابل للقيمة الحالية
// 			fmt.Print(string(memory[pointer]))
// 		case '[':
// 			// إذا كانت القيمة الحالية صفراً، انتقل إلى القوس المربع المقابل
// 			if memory[pointer] == 0 {
// 				// عدّاد للأقواس المتداخلة
// 				bracketCount := 1
				
// 				// البحث عن القوس المربع المقابل
// 				for bracketCount > 0 {
// 					i++
					
// 					// التأكد من عدم تجاوز نهاية الكود
// 					if i >= length {
// 						return
// 					}
					
// 					if source[i] == '[' {
// 						bracketCount++
// 					} else if source[i] == ']' {
// 						bracketCount--
// 					}
// 				}
// 			}
// 		case ']':
// 			// إذا كانت القيمة الحالية غير صفر، ارجع إلى القوس المربع المقابل
// 			if memory[pointer] != 0 {
// 				// عدّاد للأقواس المتداخلة
// 				bracketCount := 1
				
// 				// البحث عن القوس المربع المقابل
// 				for bracketCount > 0 {
// 					i--
					
// 					// التأكد من عدم تجاوز بداية الكود
// 					if i < 0 {
// 						return
// 					}
					
// 					if source[i] == ']' {
// 						bracketCount++
// 					} else if source[i] == '[' {
// 						bracketCount--
// 					}
// 				}
// 			}
// 		}
		
// 		i++
// 	}
// }
