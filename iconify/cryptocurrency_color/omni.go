package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Omni(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#1c347a"/><path fill="#fff" fill-rule="nonzero" d="M10.065 6.888a10.93 10.93 0 0 0-3.19 3.196V6.888zm15.004 3.11a10.93 10.93 0 0 0-3.134-3.11h3.134zm-3.088 15.084a10.933 10.933 0 0 0 3.088-3.08v3.08zM6.875 21.916a10.93 10.93 0 0 0 3.144 3.166H6.875zM26 16c0 5.514-4.486 10-10 10S6 21.514 6 16S10.486 6 16 6s10 4.486 10 10zm-10 7.292c4.02 0 7.292-3.271 7.292-7.292c0-4.02-3.271-7.292-7.292-7.292c-4.02 0-7.292 3.271-7.292 7.292c0 4.02 3.271 7.292 7.292 7.292z"/></g>`),
		g.Group(children),
	)
}