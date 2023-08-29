package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighSpeedTrain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.09 8.52c.58-.34 1.25-.52 1.92-.52H30v15H11.82c-.53 0-1.04-.21-1.42-.59L9 21H7.81a3.78 3.78 0 0 1-3.33-2H29v-1H4.086a3.825 3.825 0 0 1 1.634-4h.003l9.367-5.48Zm7.84.49a.95.95 0 0 0-.95.95v3.09c0 .52.42.95.95.95H29V9.01h-6.07ZM18.05 14c.53 0 .95-.43.95-.95V9.96a.973.973 0 0 0-.96-.96h-1.68c-.307 0-.675.212-.675.212L7.407 14H18.05Zm-5.06 3c.55 0 1-.44 1-1c0-.55-.45-1-1-1h-2c-.55 0-1 .45-1 1s.45 1 1 1h2ZM2 21h2.89c.807.626 1.818 1 2.92 1h.774l1.106 1.115l.003.002A3.006 3.006 0 0 0 11.82 24H14v6h-3v-6H7v6H4v-6H2v-3Zm15.99 9v-6h2.99v6h-2.99Zm6.99 0v-6h3v6h-3Z"/>`),
		g.Group(children),
	)
}