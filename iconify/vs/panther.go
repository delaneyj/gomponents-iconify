package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Panther(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M0 1332V232q506-12 736 6q28-43 60-77t84-69.5t130-55T1185 17q-12 229 24 299q44 14 89.5 38.5t78 46.5t78 62t67 60.5T1589 589q49 37 162 222t181 333l68 88l-58 104l-29 93q-52 90-109.5 133.5T1649 1606q38-64 20.5-119.5t-74-92T1471 1357q-65-1-118 41t-47 103t86 104t153 38l-106 53l114 18q-385 153-483 262l51-185q-76-223-235-422t-342-305q-64 16-277.5 119T0 1332zM1079 328q-16-136-9-202q-186 62-207 163q108 13 216 39zm166 386l167 217l177-25q-59-63-86-89.5t-68-53t-82-35.5t-108-14z"/>`),
		g.Group(children),
	)
}