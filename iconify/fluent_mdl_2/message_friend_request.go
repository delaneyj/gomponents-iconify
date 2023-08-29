package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFriendRequest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v768q-12-23-26-49t-31-51t-35-50t-36-42V640l-896 448l-896-448v896h960l64 128H0V384h2048zm-143 128H143l881 441l881-441zm-97 1075q55 29 99 71t76 94t48 110t17 122v64h-128v-64q0-66-25-124t-69-101t-102-69t-124-26q-66 0-124 25t-102 69t-69 102t-25 124v64h-128v-64q0-63 16-121t48-110t76-94t100-72q-54-46-83-109t-29-134q0-66 25-124t68-101t102-69t125-26q66 0 124 25t101 69t69 102t26 124q0 70-29 133t-83 110zm-400-243q0 40 15 75t41 61t61 41t75 15q40 0 75-15t61-41t41-61t15-75q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75z"/>`),
		g.Group(children),
	)
}