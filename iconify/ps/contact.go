package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 0h-64v4q-12-4-21-4H85Q68 0 55.5 12.5T43 43v64H21q-21 0-21 21t21 21h22v86H21q-21 0-21 21t21 21h22v86H21q-21 0-21 21t21 21h22v64q0 18 12.5 30.5T85 512h299q27 0 45.5-18.5T448 448v-85q27 0 45.5-18.5T512 299v-86h-4q4-12 4-21v-85h-4q4-14 4-22V43q0-18-12.5-30.5T469 0zm-21 43h21v42q0 22-21 22V43zm0 106q9 0 21-4v47q0 21-21 21v-64zm-43 299q0 21-21 21H85v-64h22q21 0 21-21t-21-21H85v-86h22q21 0 21-21t-21-21H85v-86h22q21 0 21-21t-21-21H85V43h299q21 0 21 21v384zm43-128v-64q9 0 21-4v47q0 21-21 21z"/>`),
		g.Group(children),
	)
}