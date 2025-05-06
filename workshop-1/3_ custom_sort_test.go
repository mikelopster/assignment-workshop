package customsort // ชื่อ package ควรตรงกับไฟล์ที่ผู้เรียนจะสร้าง

import (
	"reflect" // ใช้สำหรับ DeepEqual เพื่อเปรียบเทียบ slices/structs
	"testing"
)

// --- ส่วนของ Test Cases ---

func TestSortPeopleByAge_EmptySlice(t *testing.T) {
	t.Parallel() // สามารถรัน test case นี้พร้อมกับตัวอื่นได้
	input := []Person{}
	expected := []Person{}

	// เรียกฟังก์ชันที่ผู้เรียนจะสร้าง (ต้องอยู่ใน package เดียวกัน)
	SortPeopleByAge(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Test Failed: Expected %v, Got %v", expected, input)
	}
}

func TestSortPeopleByAge_SingleElement(t *testing.T) {
	t.Parallel()
	input := []Person{{"Alice", 30}}
	// คาดหวังว่า Slice เดิมจะไม่เปลี่ยนแปลง
	expected := []Person{{"Alice", 30}}

	SortPeopleByAge(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Test Failed: Expected %v, Got %v", expected, input)
	}
}

func TestSortPeopleByAge_AlreadySorted(t *testing.T) {
	t.Parallel()
	input := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}

	SortPeopleByAge(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Test Failed: Expected %v, Got %v", expected, input)
	}
}

func TestSortPeopleByAge_ReverseSorted(t *testing.T) {
	t.Parallel()
	input := []Person{
		{"Charlie", 35},
		{"Alice", 30},
		{"Bob", 25},
	}
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}

	SortPeopleByAge(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Test Failed: Expected %v, Got %v", expected, input)
	}
}

func TestSortPeopleByAge_Unsorted(t *testing.T) {
	t.Parallel()
	input := []Person{
		{"David", 28},
		{"Eve", 22},
		{"Frank", 31},
		{"Grace", 22}, // Duplicate age
	}
	expected := []Person{
		{"Eve", 22},
		{"Grace", 22},
		{"David", 28},
		{"Frank", 31},
	}
	// หมายเหตุ: ลำดับของ Eve และ Grace อาจสลับกันได้หากใช้ unstable sort
	// แต่ผลลัพธ์สุดท้ายต้องมีอายุเรียงจากน้อยไปมาก

	SortPeopleByAge(input)

	// สำหรับ test case นี้ เราอาจจะต้อง check แค่ลำดับอายุ
	// หรือถ้าต้องการให้ stable sort ก็ต้อง implement แบบนั้น
	// แบบง่ายสุดคือ DeepEqual ถ้าผู้เรียนใช้ sort.SliceStable
	if !reflect.DeepEqual(input, expected) {
		// ถ้าใช้ sort.Slice (unstable) อาจจะต้องเขียน logic check ที่ซับซ้อนกว่า
		// แต่เพื่อความง่ายของโจทย์ แนะนำให้ผู้เรียนใช้ sort.SliceStable
		// หรือยอมรับว่าลำดับของคนอายุเท่ากันอาจสลับได้
		t.Logf("Note: Order of elements with the same age might differ if using unstable sort.")
		t.Errorf("Test Failed:\nExpected (or similar order for same age): %v\nGot: %v", expected, input)

		// ทางเลือก: เช็คแค่ลำดับอายุ (ซับซ้อนขึ้นเล็กน้อย)
		// ages := make([]int, len(input))
		// for i, p := range input {
		//  ages[i] = p.Age
		// }
		// expectedAges := []int{22, 22, 28, 31}
		// if !reflect.DeepEqual(ages, expectedAges) {
		//  t.Errorf("Test Failed: Ages are not sorted correctly. Expected ages %v, Got ages %v", expectedAges, ages)
		// }
	}
}

func TestSortPeopleByAge_WithZeros(t *testing.T) {
	t.Parallel()
	input := []Person{
		{"Heidi", 30},
		{"Ivan", 0}, // Age zero
		{"Judy", 25},
	}
	expected := []Person{
		{"Ivan", 0},
		{"Judy", 25},
		{"Heidi", 30},
	}

	SortPeopleByAge(input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Test Failed: Expected %v, Got %v", expected, input)
	}
}

