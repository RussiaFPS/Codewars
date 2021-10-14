package convertStringToCamelCase

import "testing"

type dataTest struct {
	input string
	res string
	want string
}

func TestToCamelCase(t *testing.T) {
	tests:=[]dataTest{
		{input: "The_Stealth_Warrior",want:"TheStealthWarrior"},
		{input: "the-Stealth-Warrior",want: "theStealthWarrior"},
		{input: "You_have_chosen_to_translate_this_kata_For_your_convenience_we_have_provided_the_existing_test_cases_used_for_the_language_that_you_have_already_completed_as_well_as_all_of_the_other_related_fields",want: "YouHaveChosenToTranslateThisKataForYourConvenienceWeHaveProvidedTheExistingTestCasesUsedForTheLanguageThatYouHaveAlreadyCompletedAsWellAsAllOfTheOtherRelatedFields"},
	}
	for _,test:=range tests{
		test.res=ToCamelCase(test.input)
		if test.res!=test.want{
			t.Errorf("\nInput:%s\nRes:%s\nWant:%s",test.input,test.res,test.want)
		}
	}
}