package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMacaoSarChina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#186648" d="M5 17h62v38H5z"/><g stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke="#fff" d="M32.408 49.977h7m4.5-2h-16m-2.5-2h21"/><path fill="#f1b31c" stroke="#f1b31c" d="m34.467 27.023l1.66-5l1.43 4.923l-3.974-2.966l5-.124l-4.116 3.167zm8.526 1.284l-.638-3.096l2.221 2.129l-2.955-.349l2.561-1.564l-1.189 2.88zm3.101 4.74l.996-3l.859 2.954l-2.385-1.78l3-.074l-2.47 1.9zm-17.087-4.892l.638-3.096l-2.221 2.129l2.955-.349l-2.561-1.564l1.189 2.88zm-3.101 4.74l-.996-3l-.859 2.954l2.385-1.78l-3-.074l2.47 1.9z"/><g fill="#fff" stroke="#fff"><path d="M35.007 41.34c1.382 1.256 4.002 1.832 6.835 1.333s5.098-1.938 5.967-3.591c-1.382-1.256-4.003-1.833-6.836-1.333s-5.097 1.937-5.967 3.59Z"/><path d="M24.006 39.083c.87 1.653 3.135 3.09 5.967 3.59s5.454-.077 6.836-1.333c-.869-1.653-3.135-3.091-5.967-3.591s-5.453.077-6.835 1.333Z"/><path d="M35.907 29.977c-1.476 1.044-2.5 3.31-2.5 5.937s1.024 4.894 2.5 5.938c1.478-1.044 2.5-3.31 2.5-5.938s-1.022-4.893-2.498-5.937Z"/></g></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}