package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Straslines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.693 13.977a15.445 15.445 0 0 1 2.688.225a12.222 12.222 0 0 1 3.041.932a2.97 2.97 0 0 1 1.372 1.024a2.887 2.887 0 0 1 .323 1.08c.328 1.603.67 4.752.768 5.51v8.643h-1.377v1.377a1.178 1.178 0 1 1-2.35 0V31.39h-8.93v1.377a1.178 1.178 0 1 1-2.35 0V31.39H4.5v-8.643c.082-.758.425-3.887.753-5.51a2.887 2.887 0 0 1 .322-1.08a2.97 2.97 0 0 1 1.373-1.024a12.222 12.222 0 0 1 3.041-.932a15.444 15.444 0 0 1 2.704-.225Zm-5.397 3.507a.722.722 0 0 0-.692.625l-.65 4.72a.722.722 0 0 0 .594.815h12.166a.727.727 0 0 0 .727-.707a.632.632 0 0 0 0-.097l-.65-4.721a.722.722 0 0 0-.717-.625H7.321Zm-.236 8.94a1.234 1.234 0 1 0 1.229 1.24v-.01a1.229 1.229 0 0 0-1.229-1.23Zm11.265 0a1.234 1.234 0 1 0 1.234 1.234v-.005a1.224 1.224 0 0 0-1.219-1.229ZM5.877 31.391h13.631M8.407 16.015h8.541m7.009 24.37V7.615m17.923 6.106H28.734a1.332 1.332 0 0 0-1.331 1.31l-.288 16.866a1.332 1.332 0 0 0 1.332 1.355h13.721a1.332 1.332 0 0 0 1.332-1.355l-.288-16.866a1.332 1.332 0 0 0-1.332-1.31Z"/><circle cx="30.565" cy="30.459" r="1.12" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.307 13.721V11.63m-4.742 21.622v.971l-2.147 2.147m2.744-24.161l.579-.579h3.566"/><circle cx="40.05" cy="30.459" r="1.12" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.05 33.252v.971l2.147 2.147m-2.744-24.161l-.579-.579h-3.567m-3.323 9v6.199m3.137 1.157l-3.137-4.257m8.667-6.833H29.964l-.333 11.09h11.353l-.333-11.09z"/>`),
		g.Group(children),
	)
}