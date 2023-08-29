package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CmOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#007a5e" d="M0 0h170.7v512H0z"/><path fill="#ce1126" d="M170.7 0h170.6v512H170.7z"/><path fill="#fcd116" d="M341.3 0H512v512H341.3z"/><g fill="#fcd116" transform="translate(256 256) scale(5.6889)"><g id="flagCm1x10"><path id="flagCm1x11" d="M0-8L-2.5-.4L1.3.9z"/><use width="100%" height="100%" href="#flagCm1x11" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagCm1x10" transform="rotate(72)"/><use width="100%" height="100%" href="#flagCm1x10" transform="rotate(144)"/><use width="100%" height="100%" href="#flagCm1x10" transform="rotate(-144)"/><use width="100%" height="100%" href="#flagCm1x10" transform="rotate(-72)"/></g>`),
		g.Group(children),
	)
}