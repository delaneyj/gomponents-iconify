package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 43H64q-27 0-45.5 18T0 107v192q0 17 12.5 29.5T43 341h426q18 0 30.5-12.5T512 299V107q0-28-18.5-46T448 43zm-21 256v-64h-22v64h-21v-64h-21v64h-22v-64h-21v64h-43v-64q0-22-21-22t-21 22v64h-22v-64h-21v64h-21v-64h-22v64h-21v-64h-21v64H85v-64H64v64H43V107q0-22 21-22h384q21 0 21 22v192h-42zM107 107H85q-21 0-21 21v64q0 21 21 21h22q21 0 21-21v-64q0-21-21-21zm85 0h-21q-22 0-22 21v64q0 21 22 21h21q21 0 21-21v-64q0-21-21-21zm149 0h-21q-21 0-21 21v64q0 21 21 21h21q22 0 22-21v-64q0-21-22-21zm86 0h-22q-21 0-21 21v64q0 21 21 21h22q21 0 21-21v-64q0-21-21-21z"/>`),
		g.Group(children),
	)
}