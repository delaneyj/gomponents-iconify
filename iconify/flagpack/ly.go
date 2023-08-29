package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLy0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLy1)"><use href="#flagpackLy0"/><path fill="#55BA07" fill-rule="evenodd" d="M0 18h32v6H0v-6Z" clip-rule="evenodd"/><path fill="#1D1D1D" fill-rule="evenodd" d="M0 6h32v12H0V6Z" clip-rule="evenodd"/><path fill="#E11C1B" fill-rule="evenodd" d="M0 0h32v6H0V0Z" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="M15.798 15.066c-1.793-.37-2.258-1.308-2.238-2.755c0-1.53.806-3.08 2.22-3.249c1.413-.169 2.57.358 3.178 1.094c-.508-1.997-2.19-2.216-3.48-2.216c-1.944-.017-3.938 1.46-3.938 4.239c0 2.42 1.757 4.24 3.994 4.285c2.795 0 3.233-1.93 3.313-2.588a5.065 5.065 0 0 0-.466.397c-.562.517-1.178 1.084-2.583.793Zm4.465-3.775l-1.324.417l1.483.58l-.207 1.554l1.005-1.07l1.458.3l-.985-1.151l.872-1.177l-1.218.257l-.86-.972l-.224 1.262Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackLy1"><use href="#flagpackLy0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}