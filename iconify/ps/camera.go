package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 45H107V24q0-21-22-21q-21 0-21 21v21H43q-18 0-30.5 13T0 88v277q0 28 18.5 46T64 429h384q27 0 45.5-18t18.5-46V88q0-17-12.5-30T469 45zM341 88h43v43h-43V88zM43 88h256v43H43V88zm21 299q-21 0-21-22V173h185q-57 40-57 107t57 107H64zm235-22q-35 0-60.5-25T213 280t25.5-60t60.5-25t60 25t25 60t-25 60t-60 25zm170 0q0 22-21 22h-79q58-41 58-107t-58-107h100v192zm0-234h-42V88h42v43zm-170 85q-28 0-46 18.5T235 280t18 45.5t46 18.5q27 0 45.5-18.5T363 280t-18.5-45.5T299 216zm0 85q-22 0-22-21t22-21q21 0 21 21t-21 21zm-214 0h43v43H85v-43z"/>`),
		g.Group(children),
	)
}