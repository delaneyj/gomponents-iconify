package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBaby0"><g fill="none"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-width="4" d="M24 43.6c8.432 0 15.56-6.68 17.894-14.35C42.418 27.526 46 27.526 46 23.8s-3.616-3.94-4.201-5.752C39.372 10.535 32.32 4 24 4C15.675 4 8.62 10.54 6.197 18.06C5.615 19.87 2 20.01 2 23.8s3.592 3.79 4.135 5.542C8.497 36.964 15.602 43.6 24 43.6Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M41.799 18.048C39.372 10.535 32.32 4 24 4"/><path fill="#000" stroke="#000" d="M19.1 21.6c0 .826-.224 1.552-.56 2.056c-.339.508-.752.744-1.14.744c-.388 0-.8-.236-1.14-.744c-.336-.504-.56-1.23-.56-2.056c0-.827.224-1.552.56-2.056c.34-.508.752-.744 1.14-.744c.389 0 .801.236 1.14.744c.336.504.56 1.23.56 2.056Zm13.2 0c0 .826-.224 1.552-.56 2.056c-.339.508-.752.744-1.14.744c-.388 0-.801-.236-1.14-.744c-.336-.504-.56-1.23-.56-2.056c0-.827.224-1.552.56-2.056c.339-.508.752-.744 1.14-.744c.388 0 .801.236 1.14.744c.336.504.56 1.23.56 2.056Z"/><path fill="#000" fill-rule="evenodd" d="M18.498 31.75c1.93 1.3 3.768 1.95 5.511 1.95c1.742 0 3.469-.649 5.18-1.945" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M18.498 31.75c1.93 1.3 3.768 1.95 5.511 1.95c1.742 0 3.469-.649 5.18-1.945"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M31.728 6.2c.268 1.934-.321 3.347-1.769 4.239c-1.447.892-3.799 1.31-7.055 1.254"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBaby0)"/>`),
		g.Group(children),
	)
}