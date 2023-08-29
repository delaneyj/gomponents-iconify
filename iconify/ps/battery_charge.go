package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 115q0-28-18.5-46T384 51H277v42h107q21 0 21 22q0 17 13 29.5t30 12.5q21 0 21 22v42q0 22-21 22q-17 0-30 12.5T405 285q0 22-21 22h-64v42h64q27 0 45.5-18t18.5-46q27 0 45.5-18t18.5-46v-42q0-28-18.5-46T448 115zM43 285V115q0-22 21-22h85V51H64q-27 0-45.5 18T0 115v170q0 28 18.5 46T64 349h128v-42H64q-21 0-21-22zm143-96l76-170l-38-17l-111 251l171-44l-77 196l40 15l109-271z"/>`),
		g.Group(children),
	)
}