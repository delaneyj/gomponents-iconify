package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrisonSchoolBus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 149h-42q-22 0-22-21V85q0-17-12.5-29.5T363 43H43q-18 0-30.5 12.5T0 85v192q0 18 12.5 30.5T43 320h4q6 19 22.5 31t37.5 12q20 0 36.5-12.5T166 320h177q15 43 60 43q20 0 36.5-12.5T463 320h6q18 0 30.5-12.5T512 277v-85q0-17-12.5-30T469 149zM107 320q-10 0-16-6t-6-15q0-22 22-22q9 0 15 6t6 16q0 9-6 15t-15 6zm298 0q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6zm60-43q-6-18-23-30t-37-12t-36.5 11.5T346 277H166q-6-18-22.5-30T107 235q-21 0-37.5 11.5T47 277h-4V85h320v22h-22q-21 0-21 21v21q0 10 6 16t15 6h39q18 21 47 21h27q9 11 15 15v70h-4zM277 107H85q-21 0-21 21v21q0 10 6 16t15 6h192q22 0 22-22v-21q0-21-22-21z"/>`),
		g.Group(children),
	)
}