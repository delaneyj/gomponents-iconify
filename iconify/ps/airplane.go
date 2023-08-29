package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m489 230l-41-17v-42q0-43-53-43q-49 0-54 38l-23-10l2-49q0-34-21.5-70.5T245 0t-53 36.5t-21 72.5l2 49l-24 11q0-16-11.5-28.5T96 128q-53 0-53 43v51l-20 8Q0 242 0 269v77l85-20v15q0 22 22 22q9 0 15-6t6-16v-21h-9l62-13l5 86q-79 32-79 98v21h277v-21q0-71-79-101l4-89l84 19h-9v21q0 10 6 16t15 6q22 0 22-22v-12l85 19v-77q0-26-23-41zm-105-59h21v23l-21-8v-15zm-299 0h22v17l-22 13v-30zM43 294v-23l132-66l4 60zm111 175q5-12 34-32l2 32h-36zm79 0l-20-362q0-19 12-41.5T245 43t20 22t12 40l-19 364h-25zm104 0h-38l2-34q29 13 36 34zm132-175l-158-34l3-57l153 68v23h2z"/>`),
		g.Group(children),
	)
}