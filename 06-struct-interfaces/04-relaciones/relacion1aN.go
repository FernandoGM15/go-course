package main

import "fmt"

// Relación de 1 a muchos
type Video struct {
	title  string
	course Course
}

type Course struct {
	title  string
	videos []Video
}

func main() {

	video1 := Video{
		title: "01 - Introducción",
	}

	video2 := Video{
		title: "02 - Instalación",
	}

	course := Course{
		title:  "Curso de GO",
		videos: []Video{video1, video2},
	}

	video1.course = course
	video2.course = course

	// fmt.Println(video1.course.title)

	for _, v := range course.videos {
		fmt.Println(v.title)
	}
}
