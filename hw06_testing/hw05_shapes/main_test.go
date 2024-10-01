package main

import (
	"errors"
	"testing"
)

func TestSquareCircle(t *testing.T) {
	tests := []struct {
		name     string
		c        Circle
		expected float32
	}{
		{"Радиус 5", *NewCircle(5), 78.53982},
		{"Радиус 7", *NewCircle(7), 153.93805},
		{"Радиус 0", *NewCircle(0), 0},
		{"Радиус 5.5", *NewCircle(5.5), 95.03319},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.c.Square()
			if got != tc.expected {
				t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.expected)
			}
		})
	}
}

func TestSquareRectangle(t *testing.T) {
	tests := []struct {
		name     string
		c        Rectangle
		expected float32
	}{
		{"Высота 5, ширина 19", *NewRectangle(5, 19), 95},
		{"Высота 13, ширина 4", *NewRectangle(13, 4), 52},
		{"Высота 0, ширина 0", *NewRectangle(0, 0), 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.c.Square()
			if got != tc.expected {
				t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.expected)
			}
		})
	}
}

func TestSquareTriangle(t *testing.T) {
	tests := []struct {
		name     string
		c        Triangle
		expected float32
	}{
		{"Основание 5, Высота 19", *NewTriangle(5, 19), 47.5},
		{"Основание 13, Высота 4", *NewTriangle(13, 4), 26},
		{"Основание 0, Высота 0", *NewTriangle(0, 0), 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.c.Square()
			if got != tc.expected {
				t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.expected)
			}
		})
	}
}

func TestCalculateArea(t *testing.T) {
	errortarget := errors.New("переданный объект не является фигурой")
	tests := []struct {
		name     string
		c        any
		expected float32
		err      error
	}{
		{"Треугольник", NewTriangle(5, 19), 47.5, nil},
		{"Круг", NewCircle(13), 530.9292, nil},
		{"Прямоугольник", NewRectangle(6, 18), 108, nil},
		{"Кот", NewCat(), 0, errors.New("переданный объект не является фигурой")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculateArea(tc.c)
			if !errors.Is(tc.err, errortarget) {
				if errors.Is(err, errortarget) {
					t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.err)
				} else if got != tc.expected {
					t.Errorf("test(%q) = %v; want %v", tc.name, got, tc.expected)
				}
			} else if !errors.Is(tc.err, err) {
				t.Errorf("test(%q) = %v; want %v", tc.name, err, tc.err)
			}
		})
	}
}
