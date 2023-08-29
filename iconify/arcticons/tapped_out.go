package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TappedOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.858 17.726c-1 .081-3.824-3.329-5.608-1.302s2.054 4.486 5.405 5.134s7.782 3.702 7.782 10.674s-5.053 8.944-6.458 8.944s-4.27-1.162-4.27-3.945c-1.675-.054-1.883-1.153-.207-2.774s.883-7.098-.81-7.909s-4.793 1.135-5.297-1.26s4.35-6.028 4.35-6.028"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.422 31.016s.791 3.279.07 4.504c1.077 0 2.442 1.466 4.071 2.053c1.809.652 3.834.235 3.834-4.143c0-8.323-8.301-9.224-12.192-8.395c1.657.54 2.538 1.068 3.07 2.012"/><circle cx="23.275" cy="15.964" r="4.756" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.803 17.451c-2.741-3.405-.499-7.35 3.446-7.35c3.513 0 3.995 2.2 3.995 2.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.722 10.125c.163-1.254 1.5-2.375 3.622-1.7c1.283-1.486 4.093-5.107 10.687-3.54a10.006 10.006 0 0 1 6.971 13.97c-2.648 5.404-5.727 6.837-5.727 6.837"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.41 25.584c1.378-.918 5.034.115 3.377 3.08c-1.826 3.27-4.053-.186-4.053-.186"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.332 29.925C28.166 34.421 27.923 43.5 27.923 43.5m-11.709-4.42c-.154 1.177-.492 2.69-.492 2.69m12.309-18.699c1.73-2.62 3.972-4.026 4.729-3.35c1.043.93-2.313 3.472-1.296 4.134s3.433-1.525 3.433-1.525m-.813-11.5c2.054-2.81 4.199-7.475-1.035-4.724"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.682 14.235c3.002-1.82 6.295-6.782 1.633-8.485M15.092 26.761c.752.328 1.329-.573 1.329-.573s.468 2.77 2.69 1.099c.45 1.535 1.053 1.415 2.008.91c.144 1.063 1.315.815 2.01.43m-.151 6.244s-2.576 1.531-3.711-.396s2.324-2.54 2.684-1.676a1.995 1.995 0 0 1 2.447.469"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.861 37.817a5.044 5.044 0 0 0-.147-2.456"/><circle cx="25.236" cy="18.201" r=".75" fill="currentColor"/><circle cx="17.077" cy="14.739" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}