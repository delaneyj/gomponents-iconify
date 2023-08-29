package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antidote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.777 7.498a19.068 19.068 0 1 1-16.81 18.935a18.984 18.984 0 0 1 2.18-8.86"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.82 17.106a15.826 15.826 0 1 1-28.61 9.327q0-.43.023-.854"/><ellipse cx="11.093" cy="8.152" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.254" ry="3.941" transform="rotate(-33.945 11.093 8.151)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.573 4.936a3.136 3.136 0 0 1 2.763 1.123c1.309 1.944-.932 5.742-5.004 8.483s-8.434 3.386-9.742 1.442a2.503 2.503 0 0 1-.339-1.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.72 8.442c1.493 1.452-.932 5.742-5.004 8.483c-3.835 2.581-7.927 3.304-9.484 1.756m.001 6.898c6.41-3.302 10.806-.017 16.515-1.444c5.438-1.36 12.072-7.028 12.072-7.028s-6.679 13.69-18.008 13.69c-10.602 0-10.579-5.218-10.579-5.218ZM6.59 15.984l2.643 2.696M21.336 6.059l2.383 2.383M4.84 13.331l2.623 1.606c1.223.86 4.541.095 7.663-2.006s4.822-4.78 4.212-6.311L17.94 3.542"/>`),
		g.Group(children),
	)
}