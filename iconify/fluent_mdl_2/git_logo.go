package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q25 0 49 11t42 28l894 894q20 20 29 45t10 52q0 26-11 49t-28 40l-890 890q-39 39-95 39q-25 0-49-11t-42-28L39 1115q-17-17-28-41t-11-50q0-26 10-51t29-44l613-613l232 232q-12 29-12 61q0 64 46 110l8 8q2 2 5 3t6 1t10 4v567q-12 5-20 14t-18 18q-26 26-41 55t-15 67q0 34 13 63t35 52t52 35t64 13q34 0 65-11t53-31t36-50t14-65q0-22-6-44t-18-41t-28-36t-35-27V763l212 212q-12 29-12 60q0 33 12 61t34 50t49 33t62 13q33 0 61-12t50-34t33-49t13-62q0-32-12-60t-34-50t-49-34t-61-13q-27 0-52 9l-227-227q9-25 9-52q0-33-12-61t-34-50t-49-33t-62-12q-24 0-50 8L743 224L929 39q39-39 95-39z"/>`),
		g.Group(children),
	)
}