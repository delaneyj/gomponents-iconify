package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stumbleupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm171-649q0-70-50-120t-121-50t-121 50t-50 120v205q0 14-10 24t-24 10t-24-10t-10-24V478H136v102q0 71 50 121t121 50t121-50t50-121V375q0-14 10-24t24-10t24 10t10 24v35l68 68l69-68v-35zm68 103v102q0 14-10 24t-24 10t-24-10t-10-24V478l-69 68l-68-68v102q0 71 50 121t120.5 50T837 701t50-121V478H751z"/>`),
		g.Group(children),
	)
}