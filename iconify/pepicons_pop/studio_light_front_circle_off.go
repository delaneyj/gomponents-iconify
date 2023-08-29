package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightFrontCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.013 6.508C8.236 5.174 9.197 3.5 10.741 3.5h4.518c1.544 0 2.505 1.674 1.727 3.008L15.24 9.504l-1.73-1.008L15.26 5.5h-4.518l1.748 2.996l-1.728 1.008l-1.748-2.996Zm7.973 6.984c.778 1.334-.183 3.008-1.727 3.008h-4.518c-1.544 0-2.505-1.674-1.728-3.008l1.748-2.996l1.728 1.008L10.74 14.5h4.518l-1.748-2.996l1.728-1.008l1.747 2.996Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.61 13.296c-1.33.886-3.11-.067-3.11-1.665V8.37c0-1.598 1.78-2.55 3.11-1.665l2.945 1.964L11.667 10l.888 1.332l-2.946 1.964ZM10.946 10L8.5 8.369v3.262L10.947 10Zm5.444-3.296c1.33-.886 3.11.067 3.11 1.665v3.262c0 1.598-1.78 2.55-3.11 1.665l-2.945-1.964l.888-1.332l-.888-1.332l2.946-1.964ZM15.054 10l2.447 1.631V8.37L15.053 10Zm3.36 11.406a1 1 0 0 1-1.32.508L14 20.539v.961a1 1 0 1 1-2 0v-.961l-3.094 1.375a1 1 0 0 1-.812-1.828L12 18.35V16a1 1 0 1 1 2 0v2.35l3.906 1.736a1 1 0 0 1 .508 1.32Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}