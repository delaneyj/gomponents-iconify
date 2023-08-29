package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M382 382q50 11 90-22t40-83q0-41-36-74q-32-32-92-32V64q0-27-18.5-45.5T320 0q-21 0-21 21q0 22 21 22q12 0 27.5 6T363 64v21H171V64h42q0-17-12.5-30T171 21h-64q0 12 14 25t28 18v43l-23 66q-55-13-107 51q-19 29-19 53q0 45 31 76t76 31q38 0 67-24t37-61h15q33 0 49-24l32-38q-13 30-4 68q6 27 28 48.5t51 28.5zm-275-41q-28 0-46-18t-18-46q0-27 18-45.5t46-18.5h6l-28 86h81q-6 19-22.5 30.5T107 341zm42-85v-21q18 12 22 21h-22zm15-68l9-22l15 43q-23-21-24-21zm209 85q7 8 16.5 8.5T405 275q12-10 2-28q-7-12-14-34q32-7 61 28q15 18 15 38q0 30-22 48.5T395 341q-20-3-34-16.5T343 292q-6-33 13-55q6 16 17 36zm-130-15l-42-130h119zm77-45l21-42v21q-12 9-21 21z"/>`),
		g.Group(children),
	)
}