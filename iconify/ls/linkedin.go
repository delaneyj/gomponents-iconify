package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkedin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 98v498c0 47 37 84 84 84h498c23 0 45-11 59-25c15-15 25-34 25-59V98c0-47-37-84-84-84H84c-25 0-44 10-59 25C11 53 0 75 0 98zm90 66c0-32 26-60 58-60c33 0 60 28 60 60c0 33-27 59-60 59c-32 0-58-26-58-59zm161 411V270c0-7 7-13 12-13h85c12 0 12 14 12 23c24-24 55-30 87-30c78 0 128 37 128 119v206c0 7-6 13-12 13h-88c-7 0-12-7-12-13V389c0-31-9-48-44-48c-44 0-55 29-55 68v166c0 7-7 13-14 13h-87c-5 0-12-7-12-13zm-159 0V270c0-7 7-13 12-13h87c8 0 13 5 13 13v305c0 7-6 13-13 13h-87c-6 0-12-7-12-13z"/>`),
		g.Group(children),
	)
}