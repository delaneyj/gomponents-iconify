package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimeoWithCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 .4C4.698.4.4 4.698.4 10s4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6S15.302.4 10 .4zm4.401 7.75c-.508 2.916-3.348 5.387-4.201 5.951c-.854.562-1.634-.227-1.916-.824c-.324-.682-1.293-4.373-1.547-4.68c-.254-.306-1.016.307-1.016.307l-.369-.494s1.547-1.883 2.724-2.117c1.248-.25 1.246 1.951 1.546 3.174c.291 1.183.486 1.859.739 1.859c.254 0 .739-.658 1.269-1.67c.532-1.012-.022-1.906-1.061-1.27c.415-2.54 4.34-3.152 3.832-.236z"/>`),
		g.Group(children),
	)
}