package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M280 53V43q0-18-12.5-30.5T237 0h-42q-18 0-30.5 12.5T152 43v10Q86 73 44.5 129T3 256v192q0 27 18 45.5T67 512h298q28 0 46-18.5t18-45.5V256q0-71-41.5-127T280 53zm107 395q0 21-22 21H67q-22 0-22-21V256q0-70 50.5-120.5T216 85t120.5 50.5T387 256v192zm-86-192H131q-18 0-30.5 12.5T88 299v106q0 18 12.5 30.5T131 448h170q18 0 30.5-12.5T344 405V299q0-18-12.5-30.5T301 256zm0 149H131v-42h170v42zm0-85H131v-21h170v21zM195 213h42v-21h22v-43h-22v-21h-42v21h-22v43h22v21z"/>`),
		g.Group(children),
	)
}