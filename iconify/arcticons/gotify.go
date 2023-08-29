package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gotify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.632 35.406c-1.59 2.423-3.29 5.035-7.064 6.175c-3.782 1.13-9.758.967-13.407-.719c-3.658-1.695-4.99-4.922-4.31-7.422c.67-2.508 3.343-4.3 4.55-6.224c1.197-1.935.929-4.003.613-5.918a27.988 27.988 0 0 1-.613-5.095a10.325 10.325 0 0 1 1.029-3.737"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.555 12.612c-6.943-.48-1.436-6.704.958-1.916"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.181 10.13c4.55-5.266 21.4-6.74 23.076 6.666"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.822 7.273c.185-1.999 1.217-2.155 2.616-.672m15.819 10.196a1.214 1.214 0 0 1 1.312 1.082a1.345 1.345 0 0 1-2.633 0a1.217 1.217 0 0 1 1.321-1.082Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.06 18.358c-5.507 3.954 5.381 4.123 2.509-.386M28.68 11.175a5.746 5.746 0 1 1-5.745 5.746a5.744 5.744 0 0 1 5.746-5.746Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.596 15.724a1.323 1.323 0 1 1-1.197 1.312a1.257 1.257 0 0 1 1.197-1.312Zm7.901 5.424v1.785M36.29 11.294c2.227-.738 4.476 2.871 3.03 5.951m-12.392 8.209l13.14-2.864a1.398 1.398 0 0 1 1.664 1.067l.001.006l1.734 7.948a1.398 1.398 0 0 1-1.063 1.666l-13.149 2.864a1.398 1.398 0 0 1-1.665-1.066l-.001-.006l-1.733-7.949a1.398 1.398 0 0 1 1.066-1.665ZM8.884 31.378c-5.77-2.452-1.072-5.023.829-1.03M4.5 37.99q3.112 2.872 5.028-.719m21.454.497c1.676-.389 2.225.73 1.53 3.334"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.174 26.497q1.197 2.29 3.83.135"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.047 26.976l9.098 3.965l6.294-7.882M28.479 36.09l4.74-5.988m3.341-.932l6.691 3.508"/>`),
		g.Group(children),
	)
}