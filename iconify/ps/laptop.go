package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 64q0-27-18-45.5T405 0H107Q79 0 61 18.5T43 64v213H0v43q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320v-43h-43V64zM85 64q0-21 22-21h298q22 0 22 21v213h-22V85q0-21-21-21H128q-21 0-21 21v192H85V64zM64 341q-21 0-21-21h170q0 21 22 21H64zm384 0H277q22 0 22-21h170q0 21-21 21z"/>`),
		g.Group(children),
	)
}