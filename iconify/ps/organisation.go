package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organisation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M469 256v-21q0-28-18-46t-46-18H277v-43q18 0 30.5-12.5T320 85V43q0-18-12.5-30.5T277 0h-42q-18 0-30.5 12.5T192 43v42q0 18 12.5 30.5T235 128v43H107q-28 0-46 18t-18 46v21q-18 0-30.5 12.5T0 299v42q0 18 12.5 30.5T43 384h42q18 0 30.5-12.5T128 341v-42q0-18-12.5-30.5T85 256v-21q0-22 22-22h128v43q-18 0-30.5 12.5T192 299v42q0 18 12.5 30.5T235 384h42q18 0 30.5-12.5T320 341v-42q0-18-12.5-30.5T277 256v-43h128q22 0 22 22v21q-18 0-30.5 12.5T384 299v42q0 18 12.5 30.5T427 384h42q18 0 30.5-12.5T512 341v-42q0-18-12.5-30.5T469 256zM235 43h42v42h-42V43zM85 341H43v-42h42v42zm192 0h-42v-42h42v42zm150 0v-42h42v42h-42z"/>`),
		g.Group(children),
	)
}