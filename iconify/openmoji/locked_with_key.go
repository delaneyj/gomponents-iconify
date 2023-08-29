package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockedWithKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" stroke="#D0CFCE" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M14.196 35.652c-.2-7.666 5.579-16.384 12.452-16.6c6.93-.219 14.414 10.252 12.41 17.173h-3.183s.79-7.133-2.234-10.144c-.542-1.536-4.55-4.3-7.48-3.869c-1.977.291-6.392 2.807-7.31 4.581c-1.028 1.984-1.17 8.859-1.17 8.859h-3.485z"/><path fill="#FCEA2B" d="M41.698 36.361h1.63v23.372H10.283V36.361h1.63z"/><path fill="#F1B31C" d="M42.941 36.532L26.619 59.578h16.79z"/><path fill="#F4AA41" d="M55.94 32.583c3.583-1.818 6.069-5.786 6.069-10.389c0-6.321-4.69-11.446-10.473-11.446s-10.474 5.125-10.474 11.446c0 5.112 3.1 10.046 7.358 10.62"/><path fill="#F4AA41" stroke="#F4AA41" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m56.027 32.349l-.011 3.077l.156 20.725l-4.502 4.501l-3.66-4.66l2.883-2.883l-3.157-3.157l2.641-3.642l-3.14-3.14l.014-2.976l2.709-2.71l-1.241-1.241v-3.894"/><path fill="#E27022" d="m53.544 57.09l1.2-1.202l-.219-19.827l-1.214.013z"/><circle cx="51.535" cy="18.109" r="3.521" fill="#E27022" transform="rotate(-45.001 51.535 18.11)"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M39.467 33.42c.581-8.034-5.889-14.587-12.82-14.369c-6.872.217-12.45 7.083-12.45 14.37m27.501 2.9h1.63v23.371H10.283V36.321h1.63z"/><path d="M17.825 33.42c0-5.509 3.979-10.7 8.88-10.864c4.944-.165 9.559 4.79 9.144 10.864m20.176 2.606l.257 20.446l-4.561 4.56l-4.766-4.765l2.921-2.921l-3.228-3.229l3.689-3.689l-3.211-3.211v-3.03l2.745-2.745l-1.269-1.269"/><circle cx="51.535" cy="18.109" r="3.521" transform="rotate(-45.001 51.535 18.11)"/><path d="M56.017 31.636c3.57-1.769 6.027-5.59 5.993-10c-.046-6.057-4.772-10.932-10.556-10.888c-5.784.044-10.435 4.99-10.39 11.047c.038 4.898 3.134 9.022 7.376 10.399"/></g>`),
		g.Group(children),
	)
}