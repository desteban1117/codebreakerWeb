package main

import "testing"

func TestCode4X(t *testing.T){
  expected := "XXXX"
  actual := codebreaker("1234")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode3X(t *testing.T){
  expected := "XXX"
  actual := codebreaker("1235")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode2X(t *testing.T){
  expected := "XX"
  actual := codebreaker("6734")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode1X(t *testing.T){
  expected := "X"
  actual := codebreaker("7267")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode4_(t *testing.T){
  expected := "____"
  actual := codebreaker("4321")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode3_(t *testing.T){
  expected := "___"
  actual := codebreaker("4328")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode2_(t *testing.T){
  expected := "__"
  actual := codebreaker("4389")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCode1_(t *testing.T){
  expected := "_"
  actual := codebreaker("4789")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCodeXX_(t *testing.T){
  expected := "XX_"
  actual := codebreaker("1249")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}

func TestCodeX_X(t *testing.T){
  expected := "X_X"
  actual := codebreaker("1439")
  if actual != expected {
      t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
  }
}
