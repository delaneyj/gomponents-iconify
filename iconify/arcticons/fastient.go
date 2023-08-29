package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fastient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.097 34.05a16.913 16.913 0 0 0-14.599-25.6a17.677 17.677 0 0 0-5.2.8M9.2 20.55a18.538 18.538 0 0 0-.7 4.8a16.984 16.984 0 0 0 16.97 16.999h.029a17.214 17.214 0 0 0 8.199-2.1M5.5 12.55c2.4 3.6 5 8.1 11.6 6.6c2.2 2.5 15.798 17 18.898 20.4a.913.913 0 0 0 1.1.2c.1 0 .1-.1.2-.2l1.1-1.1l1.1-1.1a.91.91 0 0 0 0-1.2c-3.3-3.1-17.799-16.7-20.399-18.9c1.6-6.699-2.8-9.299-6.6-11.598"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.9 7.151s5.2 3.8 6 5.9m-9.6-2.3s3.8 5.4 5.9 6m-4.1-7.8l5.699 5.7"/>`),
		g.Group(children),
	)
}