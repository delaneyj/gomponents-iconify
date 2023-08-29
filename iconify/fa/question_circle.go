package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M896 1248v-192q0-14-9-23t-23-9H672q-14 0-23 9t-9 23v192q0 14 9 23t23 9h192q14 0 23-9t9-23zm256-672q0-88-55.5-163T958 297t-170-41q-243 0-371 213q-15 24 8 42l132 100q7 6 19 6q16 0 25-12q53-68 86-92q34-24 86-24q48 0 85.5 26t37.5 59q0 38-20 61t-68 45q-63 28-115.5 86.5T640 892v36q0 14 9 23t23 9h192q14 0 23-9t9-23q0-19 21.5-49.5T972 829q32-18 49-28.5t46-35t44.5-48t28-60.5t12.5-81zm384 192q0 209-103 385.5T1153.5 1433T768 1536t-385.5-103T103 1153.5T0 768t103-385.5T382.5 103T768 0t385.5 103T1433 382.5T1536 768z"/>`),
		g.Group(children),
	)
}