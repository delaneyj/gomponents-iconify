package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.957 11.062V3.107H9.98v-1h-.995V1c0-.553-.445-1-.996-1c-.55 0-.905.447-.905 1v1.107h-.995v1H3.023v7.955H1.085v.875l3.064.062l-.985 2.709c-.185.52-.002 1.047.516 1.232c.519.186.957-.15 1.143-.67l1.115-3.271h1.147v3c0 .553.355 1 .905 1a.997.997 0 0 0 .996-1v-3h1.089l1.064 3.334c.185.52.646.793 1.162.607c.519-.186.702-.691.518-1.213l-.956-2.729l3.033-.062v-.874h-1.939zM3.95 3.979h2.073v.938h3.894v-.938h2.132v7.084H3.95V3.979z"/>`),
		g.Group(children),
	)
}