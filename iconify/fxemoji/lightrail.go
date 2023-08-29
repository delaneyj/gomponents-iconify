package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightrail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#575A5B" d="M360 424c0 26.51-21.49 48-48 48s-48-21.49-48-48m112 0c0 26.51 21.49 48 48 48s48-21.49 48-48"/><path fill="#FF473E" d="M511.693 448L512 64H118.892L16 448z"/><path fill="#F9F9F7" d="m91.416 352l77.17-288h-49.694L41.723 352h-3.957L16 448h496v-96z"/><path fill="#F4F3EF" d="M65.693 448L168.586 64h-49.694L16 448z"/><path fill="#F4F3EF" d="m151.095 112l-8.232 48H512v-48z"/><path fill="#132028" d="M512 496H0v-32h512v32zM89.416 352l77.169-288h-49.693L39.723 352h49.693zM512 160H181.069l-36.563 170.937l-10.895 40.661L206.746 352H256v88h112v-88h144V160zM360 424h-96v-32h96v32z"/><path fill="#FF473E" d="M134.208 371.438L216.586 64h-49.694L89.723 352H39.416l-12.861 48h1.063z"/><path fill="#FFD469" d="M37.212 360.409c6.402 1.715 10.201 8.295 8.485 14.697c-1.715 6.402-8.295 10.201-14.697 8.485"/><path fill="#575A5B" d="M256 352h-4V160h4v192zm116-192h-4v192h4V160z"/>`),
		g.Group(children),
	)
}