package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MhFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#3b5aa3" d="M0 0h639.9v480H0z"/><path fill="#e2ae57" d="M0 467L639.9 0v87L0 480v-13z"/><path fill="#fff" d="M22.4 480L640 179.2l-.1-95.5L0 480h22.4zm153-464.8L169 118l-27-65.6l10.4 69.8l-41.9-56.4l27.5 64.3l-55-42.6l42.8 53.6l-62.1-27.6l54.4 41.2l-67.7-9l64 25.4L14 180.3l100.6 6.7l-63.7 26.2l67-9l-54.3 40l63-27.6l-43 54l54.6-41.3l-27 62.9l43.6-54.7l-11.8 68.1l27.5-63.7l6.2 100.7l9.7-100.4l23.7 64l-9-69l43.4 54.8l-28.6-64l54.6 44l-43.4-54.9l64.9 27l-57.4-41.9l69.9 11.8l-67-25.7l104.1-6.5l-104-9.7l68.5-22.8l-71 9l58.6-41l-66 26.5l45.6-55.3l-55.6 43.4l26.7-66.4l-43.1 56.4l9.3-70.4l-25.7 66.5l-9.6-102.8z"/></g>`),
		g.Group(children),
	)
}