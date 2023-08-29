package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hryvnia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 448H676q-41 34-100 64q-50 28-107 64h363q26 0 45 18.5t19 45t-19 45.5t-45 19H298q-42 41-42 64q0 53 37.5 90.5T384 896q85 0 147.5-10t88.5-22t48-22t36-10q26 0 45 18.5t19 45.5q0 34-29 53l-11 12q-75 59-345 63h-31q-48 0-87.5-13.5t-70.5-42t-48.5-79.5T128 768q0-30 9-64H64q-26 0-45-18.5t-19-45T19 595t45-19h156q41-34 100-64q50-28 107-64H64q-26 0-45-19T0 383.5t19-45T64 320h534q42-41 42-64q0-53-37.5-90.5T512 128q-85 0-147.5 10T276 160t-48 22t-36 10q-26 0-45-19t-19-45q0-34 29-53q1-3 5.5-8T192 49t63-23T369 8t175-8q48 0 87.5 13.5t70.5 42t48.5 79.5T768 256q0 31-9 64h73q26 0 45 18.5t19 45t-19 45.5t-45 19z"/>`),
		g.Group(children),
	)
}