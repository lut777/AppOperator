package v1

import (
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"testing"
)

var usrnameTests = []struct{
	 give string
	 want bool
 }{
	{	give: "", 			want: true},
	{	give: "aabb", 		want: false},
	{   give: "a1a_b2b", 	want: true},
	{   give: "aa_aa_bb", 	want: false},
	{   give: "aa#a_bb", 	want: false},
	{   give: "aaa_b^b", 	want: false},
 }
func TestCcheckUsrname(t *testing.T){
	for _, unit := range usrnameTests {
		t.Run(unit.give, func(t *testing.T) {
			result := checkUsrname(unit.give)
			assert.Equal(t, unit.want, result)
		})
	}
}


var typeTests = []struct{
	give string
	want bool
}{
	{	give: "full", 	 want: true},
	{   give: "partial", want: true},
	{   give: "", 		 want: true},
	{   give: "wtf", 	 want: false},
}
func TestCheckType(t *testing.T){
	for _, unit := range typeTests {
		t.Run(unit.give, func(t *testing.T) {
			result := checkType(unit.give)
			assert.Equal(t, unit.want, result)
		})
	}
}


var maxBackupsTests = []struct{
	give []int32
	want bool
}{
	{	give: []int32{0},  want: false},
	{   give: []int32{1},  want: true},
	{   give: []int32{8},  want: true},
	{   give: []int32{10}, want: true},
	{   give: []int32{11}, want: false},
}
func TestCheckMaxB(t *testing.T){
	for _, unit := range maxBackupsTests {
		t.Run("", func(t *testing.T) {
			result := checkMaxBackups(unit.give[0])
			assert.Equal(t, unit.want, result)
		})
	}
}


var addressTest = []struct{
	give string
	want bool
}{
	{	give: "abc.aaa", 				want: true},
	{   give: "http://www.baidu.com", 	want: true},
	{   give: "https://www.google.com", want: true},
	{   give: "abc.aaa.ccc/123", 		want: true},
	{   give: "abc.aaa.ccc/123/ijk",	want: true},
	{   give: "abc/aaa", 				want: false},
}

func TestCAddress(t *testing.T){
	for _, unit := range addressTest {
		t.Run(unit.give, func(t *testing.T) {
			result := checkAddress(unit.give)
			assert.Equal(t, unit.want, result)
		})
	}
}


var m500, _ = resource.ParseQuantity("500M")
var m600, _ = resource.ParseQuantity("600M")
var resourceRequirementsTest = []struct{
	give []corev1.ResourceRequirements
	want bool
}{
	{	give: []corev1.ResourceRequirements{
		{Limits: map[corev1.ResourceName]resource.Quantity{},
			Requests: map[corev1.ResourceName]resource.Quantity{}},
	}, 				want: false},
	{	give: []corev1.ResourceRequirements{
			{Limits: map[corev1.ResourceName]resource.Quantity{"memory": m500},
		Requests: map[corev1.ResourceName]resource.Quantity{}},
	}, 				want: false},
	{	give: []corev1.ResourceRequirements{
			{Limits: map[corev1.ResourceName]resource.Quantity{},
			Requests: map[corev1.ResourceName]resource.Quantity{"memory": m500}},
	}, 				want: true},
	{	give: []corev1.ResourceRequirements{
			{Limits: map[corev1.ResourceName]resource.Quantity{"memory": m600},
			Requests: map[corev1.ResourceName]resource.Quantity{"memory": m500}},
	}, 				want: true},
	{	give: []corev1.ResourceRequirements{
		{Limits: map[corev1.ResourceName]resource.Quantity{"memory": m500},
			Requests: map[corev1.ResourceName]resource.Quantity{"memory": m500}},
	}, 				want: true},
	{	give: []corev1.ResourceRequirements{
			{Limits: map[corev1.ResourceName]resource.Quantity{"memory": m500},
			Requests: map[corev1.ResourceName]resource.Quantity{"memory": m600}},
	}, 				want: false},
}
func TestResourceRequirements(t *testing.T){
	for _, unit := range resourceRequirementsTest {
		t.Run("", func(t *testing.T) {
			result, _ := checkResourceRequirements(unit.give[0])
			assert.Equal(t, unit.want, result)
		})
	}
}


var storageSizeTest = []struct{
	give string
	want bool
}{
	{	give: "500M", 	 want: true},
	{   give: "600M", 	 want: true},
	{   give: "", 		 want: true},
	{   give: "400M", 	 want: false},
	{   give: "400A", 	 want: false},
	{   give: "400", 	 want: false},
}
func TestStorageSize(t *testing.T){
	for _, unit := range storageSizeTest {
		t.Run("", func(t *testing.T) {
			result, _ := checkStorageSize(unit.give, m500)
			assert.Equal(t, unit.want, result)
		})
	}
}