package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightFrontOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.013 3.508C5.236 2.174 6.197.5 7.741.5h4.518c1.543 0 2.505 1.674 1.727 3.008L12.24 6.504l-1.73-1.008L12.26 2.5H7.741l1.748 2.996L7.76 6.504L6.013 3.508Zm7.973 6.984c.778 1.334-.184 3.008-1.727 3.008H7.741c-1.544 0-2.505-1.674-1.728-3.008l1.748-2.996L9.49 8.504L7.74 11.5h4.518L10.51 8.504l1.728-1.008l1.747 2.996Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.61 10.296c-1.33.886-3.11-.067-3.11-1.665V5.37c0-1.598 1.78-2.55 3.11-1.665l2.945 1.964L8.667 7l.888 1.332l-2.946 1.964ZM7.946 7L5.5 5.369V8.63L7.947 7Zm5.444-3.296c1.33-.886 3.11.067 3.11 1.665V8.63c0 1.598-1.78 2.55-3.11 1.665l-2.945-1.964L11.333 7l-.888-1.332l2.946-1.964ZM12.054 7L14.5 8.631V5.37L12.053 7Zm3.36 11.406a1 1 0 0 1-1.32.508L11 17.539v.961a1 1 0 1 1-2 0v-.961l-3.094 1.375a1 1 0 0 1-.812-1.828L9 15.35V13a1 1 0 1 1 2 0v2.35l3.906 1.736a1 1 0 0 1 .508 1.32Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}