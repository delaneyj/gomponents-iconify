package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M155 304H48V176h64q27 0 45.5-19t18.5-45V48h64v85l43-42V27q0-8-7-15t-15-7H114L5 112v213q0 8 7 15t15 7h128v-43zM133 48v64q0 8-6.5 14.5T112 133H48zm320 85H306l-23 22l-43 42l-43 43v213q0 8 7 15t15 7h234q8 0 15-7t7-15V155q0-8-7-15t-15-7zm-170 86l42-43v64q0 8-6.5 14.5T304 261h-64zm149 213H240V304h64q27 0 45.5-19t18.5-45v-64h64v256z"/>`),
		g.Group(children),
	)
}