package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroBill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 21H43q-18 0-30.5 13T0 64v256q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V64q0-17-12.5-30T469 21zM149 320H43V64h106v256zm128 0h-21v-75q0-10 11-10q10 0 10 10v75zm128 0h-21v-75q0-10 11-10q10 0 10 10v75zm64 0h-21v-75q0-23-15-38t-38-15t-38.5 15.5T341 245v75h-21v-75q0-23-15-38t-38-15t-38.5 15.5T213 245v75h-21V192h277v128zm0-171H192V64h277v85zM85 299q22 0 22-22v-42q0-22-22-22q-9 0-15 6t-6 16v42q0 10 6 16t15 6zm278-171h64q9 0 15-6t6-15q0-10-6-16t-15-6h-64q-22 0-22 22q0 9 6 15t16 6z"/>`),
		g.Group(children),
	)
}