package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoStudioCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPhotoStudioCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.4 5.5a1 1 0 0 1 1-1h6.2a1 1 0 0 1 1 1v3.8a1 1 0 0 1-1 1H5.4a1 1 0 0 1-1-1V5.5Zm7.2 0H5.4v3.8h6.2V5.5Z"/><path d="M3.65 5a.5.5 0 0 1 .5-.5h8.7a.5.5 0 0 1 0 1h-8.7a.5.5 0 0 1-.5-.5Zm1.74 4.598L4.61 13.5h7.78l-.78-3.902l.98-.196l.78 3.902a1 1 0 0 1-.98 1.196H4.61a1 1 0 0 1-.98-1.196l.78-3.902l.98.196ZM22.25 20.5a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 .5-.5ZM16.75 10h.604c.137 0 .235.026.302.06c.359.182.924.508 1.385.944c.473.446.709.881.709 1.29c0 .41-.236.845-.71 1.29c-.46.436-1.025.763-1.384.944a.668.668 0 0 1-.302.06h-.604a1 1 0 0 1-1-1V11a1 1 0 0 1 1-1Zm1.358 5.42a1.66 1.66 0 0 1-.754.168h-.604a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h.604c.262 0 .52.05.754.168c.795.402 2.642 1.51 2.642 3.126c0 1.615-1.847 2.724-2.642 3.126Z"/><path d="M20.25 11.7a.5.5 0 0 0-.5.5v7.565a.5.5 0 1 0 1 0V12.2a.5.5 0 0 0-.5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPhotoStudioCircleFilled0)"/></g>`),
		g.Group(children),
	)
}