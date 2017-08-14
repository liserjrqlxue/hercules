package hercules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMailmap(t *testing.T) {
	contents := `#
# This list is used by git-shortlog to fix a few botched name translations
# in the git archive, either because the author's full name was messed up
# and/or not always written the same way, making contributions from the
# same person appearing not to be so.
#

<nico@fluxnic.net> <nico@cam.org>
Alejandro R. Sedeño <asedeno@MIT.EDU> <asedeno@mit.edu>
Alex Bennée <kernel-hacker@bennee.com>
Alex Riesen <raa.lkml@gmail.com> <fork0@t-online.de>
Alex Riesen <raa.lkml@gmail.com> <raa@limbo.localdomain>
Alex Riesen <raa.lkml@gmail.com> <raa@steel.home>
Alex Vandiver <alex@chmrr.net> <alexmv@MIT.EDU>
Alexander Gavrilov <angavrilov@gmail.com>
Alexander Kuleshov <kuleshovmail@gmail.com>
Alexey Shumkin <alex.crezoff@gmail.com> <zapped@mail.ru>
Alexey Shumkin <alex.crezoff@gmail.com> <Alex.Crezoff@gmail.com>
Anders Kaseorg <andersk@MIT.EDU> <andersk@ksplice.com>
Anders Kaseorg <andersk@MIT.EDU> <andersk@mit.edu>
Aneesh Kumar K.V <aneesh.kumar@gmail.com>
Amos Waterland <apw@debian.org> <apw@rossby.metr.ou.edu>
Amos Waterland <apw@debian.org> <apw@us.ibm.com>
Ben Walton <bdwalton@gmail.com> <bwalton@artsci.utoronto.ca>
Benoit Sigoure <tsunanet@gmail.com> <tsuna@lrde.epita.fr>
Bernt Hansen <bernt@norang.ca> <bernt@alumni.uwaterloo.ca>
Brandon Casey <drafnel@gmail.com> <casey@nrlssc.navy.mil>
brian m. carlson <sandals@crustytoothpaste.ath.cx> Brian M. Carlson <sandals@crustytoothpaste.ath.cx>
brian m. carlson <sandals@crustytoothpaste.ath.cx> <sandals@crustytoothpaste.net>
Bryan Larsen <bryan@larsen.st> <bryan.larsen@gmail.com>
Bryan Larsen <bryan@larsen.st> <bryanlarsen@yahoo.com>
Cheng Renquan <crquan@gmail.com>
Chris Shoemaker <c.shoemaker@cox.net>
Chris Wright <chrisw@sous-sol.org> <chrisw@osdl.org>
Cord Seele <cowose@gmail.com> <cowose@googlemail.com>
Christian Couder <chriscool@tuxfamily.org> <christian.couder@gmail.com>
Christian Stimming <stimming@tuhh.de> <chs@ckiste.goetheallee>
Csaba Henk <csaba@gluster.com> <csaba@lowlife.hu>
Dan Johnson <computerdruid@gmail.com>
Dana L. How <danahow@gmail.com> <how@deathvalley.cswitch.com>
Dana L. How <danahow@gmail.com> Dana How
Daniel Barkalow <barkalow@iabervon.org>
Daniel Trstenjak <daniel.trstenjak@gmail.com> <daniel.trstenjak@online.de>
Daniel Trstenjak <daniel.trstenjak@gmail.com> <trsten@science-computing.de>
David Brown <git@davidb.org> <davidb@quicinc.com>
David D. Kilzer <ddkilzer@kilzer.net>
David Kågedal <davidk@lysator.liu.se>
David Reiss <dreiss@facebook.com> <dreiss@dreiss-vmware.(none)>
David S. Miller <davem@davemloft.net>
David Turner <novalis@novalis.org> <dturner@twopensource.com>
David Turner <novalis@novalis.org> <dturner@twosigma.com>
Deskin Miller <deskinm@umich.edu>
Dirk Süsserott <newsletter@dirk.my1.cc>
Eric Blake <eblake@redhat.com> <ebb9@byu.net>
Eric Hanchrow <eric.hanchrow@gmail.com> <offby1@blarg.net>
Eric S. Raymond <esr@thyrsus.com>
Eric Wong <e@80x24.org> <normalperson@yhbt.net>
Erik Faye-Lund <kusmabite@gmail.com> <kusmabite@googlemail.com>
Eyvind Bernhardsen <eyvind.bernhardsen@gmail.com> <eyvind-git@orakel.ntnu.no>
Florian Achleitner <florian.achleitner.2.6.31@gmail.com> <florian.achleitner2.6.31@gmail.com>
Franck Bui-Huu <vagabon.xyz@gmail.com> <fbuihuu@gmail.com>
Frank Lichtenheld <frank@lichtenheld.de> <djpig@debian.org>
Frank Lichtenheld <frank@lichtenheld.de> <flichtenheld@astaro.com>
Fredrik Kuivinen <frekui@gmail.com> <freku045@student.liu.se>
Frédéric Heitzmann <frederic.heitzmann@gmail.com>
Garry Dolley <gdolley@ucla.edu> <gdolley@arpnetworks.com>
Greg Price <price@mit.edu> <price@MIT.EDU>
Greg Price <price@mit.edu> <price@ksplice.com>
Heiko Voigt <hvoigt@hvoigt.net> <git-list@hvoigt.net>
H. Merijn Brand <h.m.brand@xs4all.nl> H.Merijn Brand <h.m.brand@xs4all.nl>
H. Peter Anvin <hpa@zytor.com> <hpa@bonde.sc.orionmulti.com>
H. Peter Anvin <hpa@zytor.com> <hpa@smyrno.hos.anvin.org>
H. Peter Anvin <hpa@zytor.com> <hpa@tazenda.sc.orionmulti.com>
H. Peter Anvin <hpa@zytor.com> <hpa@trantor.hos.anvin.org>
Han-Wen Nienhuys <hanwen@google.com> Han-Wen Nienhuys <hanwen@xs4all.nl>
Horst H. von Brand <vonbrand@inf.utfsm.cl>
J. Bruce Fields <bfields@citi.umich.edu> <bfields@fieldses.org>
J. Bruce Fields <bfields@citi.umich.edu> <bfields@pig.linuxdev.us.dell.com>
J. Bruce Fields <bfields@citi.umich.edu> <bfields@puzzle.fieldses.org>
Jakub Narębski <jnareb@gmail.com>
James Y Knight <jknight@itasoftware.com> <foom@fuhm.net>
# The 2 following authors are probably the same person,
# but both emails bounce.
Jason McMullan <jason.mcmullan@timesys.com>
Jason McMullan <mcmullan@netapp.com>
Jason Riedy <ejr@eecs.berkeley.edu> <ejr@EECS.Berkeley.EDU>
Jason Riedy <ejr@eecs.berkeley.edu> <ejr@cs.berkeley.edu>
Jay Soffian <jaysoffian@gmail.com> <jaysoffian+git@gmail.com>
Jeff King <peff@peff.net> <peff@github.com>
Jeff Muizelaar <jmuizelaar@mozilla.com> <jeff@infidigm.net>
Jens Axboe <axboe@kernel.dk> <axboe@suse.de>
Jens Axboe <axboe@kernel.dk> <jens.axboe@oracle.com>
Jens Lindström <jl@opera.com> Jens Lindstrom <jl@opera.com>
Jim Meyering <jim@meyering.net> <meyering@redhat.com>
Joachim Berdal Haga <cjhaga@fys.uio.no>
Johannes Schindelin <Johannes.Schindelin@gmx.de> <johannes.schindelin@gmx.de>
Johannes Sixt <j6t@kdbg.org> <J.Sixt@eudaptics.com>
Johannes Sixt <j6t@kdbg.org> <j.sixt@viscovery.net>
Johannes Sixt <j6t@kdbg.org> <johannes.sixt@telecom.at>
John 'Warthog9' Hawley <warthog9@kernel.org> <warthog9@eaglescrag.net>
Jon Loeliger <jdl@jdl.com> <jdl@freescale.com>
Jon Loeliger <jdl@jdl.com> <jdl@freescale.org>
Jon Seymour <jon.seymour@gmail.com> <jon@blackcubes.dyndns.org>
Jonathan Nieder <jrnieder@gmail.com> <jrnieder@uchicago.edu>
Jonathan del Strother <jon.delStrother@bestbefore.tv> <maillist@steelskies.com>
Josh Triplett <josh@joshtriplett.org> <josh@freedesktop.org>
Josh Triplett <josh@joshtriplett.org> <josht@us.ibm.com>
Julian Phillips <julian@quantumfyre.co.uk> <jp3@quantumfyre.co.uk>
Junio C Hamano <gitster@pobox.com> <gitster@pobox.com>
Junio C Hamano <gitster@pobox.com> <junio@hera.kernel.org>
Junio C Hamano <gitster@pobox.com> <junio@kernel.org>
Junio C Hamano <gitster@pobox.com> <junio@pobox.com>
Junio C Hamano <gitster@pobox.com> <junio@twinsun.com>
Junio C Hamano <gitster@pobox.com> <junkio@cox.net>
Junio C Hamano <gitster@pobox.com> <junkio@twinsun.com>
Karl Wiberg <kha@treskal.com> Karl  Hasselström
Karl Wiberg <kha@treskal.com> <kha@yoghurt.hemma.treskal.com>
Karsten Blees <blees@dcon.de> <karsten.blees@dcon.de>
Karsten Blees <blees@dcon.de> <karsten.blees@gmail.com>
Kay Sievers <kay.sievers@vrfy.org> <kay.sievers@suse.de>
Kay Sievers <kay.sievers@vrfy.org> <kay@mam.(none)>
Kazuki Saitoh <ksaitoh560@gmail.com> kazuki saitoh <ksaitoh560@gmail.com>
Keith Cascio <keith@CS.UCLA.EDU> <keith@cs.ucla.edu>
Kent Engstrom <kent@lysator.liu.se>
Kevin Leung <kevinlsk@gmail.com>
Kirill Smelkov <kirr@navytux.spb.ru> <kirr@landau.phys.spbu.ru>
Kirill Smelkov <kirr@navytux.spb.ru> <kirr@mns.spb.ru>
Knut Franke <Knut.Franke@gmx.de> <k.franke@science-computing.de>
Lars Doelle <lars.doelle@on-line ! de>
Lars Doelle <lars.doelle@on-line.de>
Lars Noschinski <lars@public.noschinski.de> <lars.noschinski@rwth-aachen.de>
Li Hong <leehong@pku.edu.cn>
Linus Torvalds <torvalds@linux-foundation.org> <torvalds@evo.osdl.org>
Linus Torvalds <torvalds@linux-foundation.org> <torvalds@g5.osdl.org>
Linus Torvalds <torvalds@linux-foundation.org> <torvalds@osdl.org>
Linus Torvalds <torvalds@linux-foundation.org> <torvalds@ppc970.osdl.org.(none)>
Linus Torvalds <torvalds@linux-foundation.org> <torvalds@ppc970.osdl.org>
Linus Torvalds <torvalds@linux-foundation.org> <torvalds@woody.linux-foundation.org>
Lukas Sandström <luksan@gmail.com> <lukass@etek.chalmers.se>
Marc Khouzam <marc.khouzam@ericsson.com> <marc.khouzam@gmail.com>
Marc-André Lureau <marcandre.lureau@gmail.com>
Marco Costalba <mcostalba@gmail.com> <mcostalba@yahoo.it>
Mark Levedahl <mdl123@verizon.net> <mlevedahl@gmail.com>
Mark Rada <marada@uwaterloo.ca>
Martin Langhoff <martin@laptop.org> <martin@catalyst.net.nz>
Martin von Zweigbergk <martinvonz@gmail.com> <martin.von.zweigbergk@gmail.com>
Matt Draisey <matt@draisey.ca> <mattdraisey@sympatico.ca>
Matt Kraai <kraai@ftbfs.org> <matt.kraai@amo.abbott.com>
Matt McCutchen <matt@mattmccutchen.net> <hashproduct@gmail.com>
Matthias Kestenholz <matthias@spinlock.ch> <mk@spinlock.ch>
Matthias Urlichs <matthias@urlichs.de> <smurf@kiste.(none)>
Matthias Urlichs <matthias@urlichs.de> <smurf@smurf.noris.de>
Michael Coleman <tutufan@gmail.com>
Michael J Gruber <git@grubix.eu> <michaeljgruber+gmane@fastmail.fm>
Michael J Gruber <git@grubix.eu> <git@drmicha.warpmail.net>
Michael S. Tsirkin <mst@kernel.org> <mst@redhat.com>
Michael S. Tsirkin <mst@kernel.org> <mst@mellanox.co.il>
Michael S. Tsirkin <mst@kernel.org> <mst@dev.mellanox.co.il>
Michael W. Olson <mwolson@gnu.org>
Michael Witten <mfwitten@gmail.com> <mfwitten@MIT.EDU>
Michael Witten <mfwitten@gmail.com> <mfwitten@mit.edu>
Michal Rokos <michal.rokos@nextsoft.cz> <rokos@nextsoft.cz>
Michele Ballabio <barra_cuda@katamail.com>
Miklos Vajna <vmiklos@frugalware.org> <vmiklos@suse.cz>
Namhyung Kim <namhyung@gmail.com> <namhyung.kim@lge.com>
Namhyung Kim <namhyung@gmail.com> <namhyung@kernel.org>
Nanako Shiraishi <nanako3@lavabit.com> <nanako3@bluebottle.com>
Nanako Shiraishi <nanako3@lavabit.com>
Nelson Elhage <nelhage@mit.edu> <nelhage@MIT.EDU>
Nelson Elhage <nelhage@mit.edu> <nelhage@ksplice.com>
Nguyễn Thái Ngọc Duy <pclouds@gmail.com>
Nick Stokoe <nick@noodlefactory.co.uk> Nick Woolley <nick@noodlefactory.co.uk>
Nick Stokoe <nick@noodlefactory.co.uk> Nick Woolley <nickwoolley@yahoo.co.uk>
Nicolas Morey-Chaisemartin <devel-git@morey-chaisemartin.com> <nicolas.morey@free.fr>
Nicolas Morey-Chaisemartin <devel-git@morey-chaisemartin.com> <nmorey@kalray.eu>
Nicolas Sebrecht <nicolas.s.dev@gmx.fr> <ni.s@laposte.net>
Paolo Bonzini <bonzini@gnu.org> <paolo.bonzini@lu.unisi.ch>
Pascal Obry <pascal@obry.net> <pascal.obry@gmail.com>
Pascal Obry <pascal@obry.net> <pascal.obry@wanadoo.fr>
Pat Notz <patnotz@gmail.com> <pknotz@sandia.gov>
Patrick Steinhardt <ps@pks.im> <patrick.steinhardt@elego.de>
Paul Mackerras <paulus@samba.org> <paulus@dorrigo.(none)>
Paul Mackerras <paulus@samba.org> <paulus@pogo.(none)>
Peter Baumann <waste.manager@gmx.de> <Peter.B.Baumann@stud.informatik.uni-erlangen.de>
Peter Baumann <waste.manager@gmx.de> <siprbaum@stud.informatik.uni-erlangen.de>
Peter Krefting <peter@softwolves.pp.se> <peter@softwolves.pp.se>
Peter Krefting <peter@softwolves.pp.se> <peter@svarten.intern.softwolves.pp.se>
Petr Baudis <pasky@ucw.cz> <pasky@suse.cz>
Petr Baudis <pasky@ucw.cz> <xpasky@machine>
Phil Hord <hordp@cisco.com> <phil.hord@gmail.com>
Philip Jägenstedt <philip@foolip.org> <philip.jagenstedt@gmail.com>
Philipp A. Hartmann <pah@qo.cx> <ph@sorgh.de>
Philippe Bruhat <book@cpan.org>
Ralf Thielow <ralf.thielow@gmail.com> <ralf.thielow@googlemail.com>
Ramsay Jones <ramsay@ramsayjones.plus.com> <ramsay@ramsay1.demon.co.uk>
René Scharfe <l.s.r@web.de> <rene.scharfe@lsrfire.ath.cx>
Richard Hansen <rhansen@rhansen.org> <hansenr@google.com>
Richard Hansen <rhansen@rhansen.org> <rhansen@bbn.com>
Robert Fitzsimons <robfitz@273k.net>
Robert Shearman <robertshearman@gmail.com> <rob@codeweavers.com>
Robert Zeh <robert.a.zeh@gmail.com>
Robin Rosenberg <robin.rosenberg@dewire.com> <robin.rosenberg.lists@dewire.com>
Rutger Nijlunsing <rutger.nijlunsing@gmail.com> <rutger@nospam.com>
Rutger Nijlunsing <rutger.nijlunsing@gmail.com> <git@tux.tmfweb.nl>
Ryan Anderson <ryan@michonline.com> <rda@google.com>
Salikh Zakirov <salikh.zakirov@gmail.com> <Salikh.Zakirov@Intel.com>
Sam Vilain <sam@vilain.net> <sam.vilain@catalyst.net.nz>
Sam Vilain <sam@vilain.net> sam@vilain.net
Santi Béjar <santi@agolina.net> <sbejar@gmail.com>
Sean Estabrooks <seanlkml@sympatico.ca>
Sebastian Schuberth <sschuberth@gmail.com> <sschuberth@visageimaging.com>
Seth Falcon <seth@userprimary.net> <sfalcon@fhcrc.org>
Shawn O. Pearce <spearce@spearce.org>
Simon Hausmann <hausmann@kde.org> <simon@lst.de>
Simon Hausmann <hausmann@kde.org> <shausman@trolltech.com>
Stefan Beller <stefanbeller@gmail.com> <stefanbeller@googlemail.com>
Stefan Beller <stefanbeller@gmail.com> <sbeller@google.com>
Stefan Naewe <stefan.naewe@gmail.com> <stefan.naewe@atlas-elektronik.com>
Stefan Naewe <stefan.naewe@gmail.com> <stefan.naewe@googlemail.com>
Stefan Sperling <stsp@elego.de> <stsp@stsp.name>
Štěpán Němec <stepnem@gmail.com> <stepan.nemec@gmail.com>
Stephen Boyd <bebarino@gmail.com> <sboyd@codeaurora.org>
Steven Drake <sdrake@xnet.co.nz> <sdrake@ihug.co.nz>
Steven Grimm <koreth@midwinter.com> <sgrimm@sgrimm-mbp.local>
Steven Grimm <koreth@midwinter.com> koreth@midwinter.com
Steven Walter <stevenrwalter@gmail.com> <swalter@lexmark.com>
Steven Walter <stevenrwalter@gmail.com> <swalter@lpdev.prtdev.lexmark.com>
Sven Verdoolaege <skimo@kotnet.org> <Sven.Verdoolaege@cs.kuleuven.ac.be>
Sven Verdoolaege <skimo@kotnet.org> <skimo@liacs.nl>
SZEDER Gábor <szeder.dev@gmail.com> <szeder@ira.uka.de>
Tay Ray Chuan <rctay89@gmail.com>
Ted Percival <ted@midg3t.net> <ted.percival@quest.com>
Theodore Ts'o <tytso@mit.edu>
Thomas Ackermann <th.acker@arcor.de> <th.acker66@arcor.de>
Thomas Rast <tr@thomasrast.ch> <trast@student.ethz.ch>
Thomas Rast <tr@thomasrast.ch> <trast@inf.ethz.ch>
Thomas Rast <tr@thomasrast.ch> <trast@google.com>
Timo Hirvonen <tihirvon@gmail.com> <tihirvon@ee.oulu.fi>
Toby Allsopp <Toby.Allsopp@navman.co.nz> <toby.allsopp@navman.co.nz>
Tom Grennan <tmgrennan@gmail.com> <tgrennan@redback.com>
Tommi Virtanen <tv@debian.org> <tv@eagain.net>
Tommi Virtanen <tv@debian.org> <tv@inoi.fi>
Tommy Thorn <tommy-git@thorn.ws> <tt1729@yahoo.com>
Tony Luck <tony.luck@intel.com>
Tor Arne Vestbø <torarnv@gmail.com> <tavestbo@trolltech.com>
Trần Ngọc Quân <vnwildman@gmail.com> Tran Ngoc Quan <vnwildman@gmail.com>
Trent Piepho <tpiepho@gmail.com> <tpiepho@freescale.com>
Trent Piepho <tpiepho@gmail.com> <xyzzy@speakeasy.org>
Uwe Kleine-König <u.kleine-koenig@pengutronix.de> <Uwe.Kleine-Koenig@digi.com>
Uwe Kleine-König <u.kleine-koenig@pengutronix.de> <ukleinek@informatik.uni-freiburg.de>
Uwe Kleine-König <u.kleine-koenig@pengutronix.de> <uzeisberger@io.fsforth.de>
Uwe Kleine-König <u.kleine-koenig@pengutronix.de> <zeisberg@informatik.uni-freiburg.de>
Ville Skyttä <ville.skytta@iki.fi> <scop@xemacs.org>
Vitaly "_Vi" Shukela <public_vi@tut.by>
W. Trevor King <wking@tremily.us> <wking@drexel.edu>
William Pursell <bill.pursell@gmail.com>
YONETANI Tomokazu <y0n3t4n1@gmail.com> <qhwt+git@les.ath.cx>
YONETANI Tomokazu <y0n3t4n1@gmail.com> <y0netan1@dragonflybsd.org>
YOSHIFUJI Hideaki <yoshfuji@linux-ipv6.org>
# the two anonymous contributors are different persons:
anonymous <linux@horizon.com>
anonymous <linux@horizon.net>
İsmail Dönmez <ismail@pardus.org.tr>`
	mm := ParseMailmap(contents)
	assert.Equal(t, mm["ismail@pardus.org.tr"].Name, "İsmail Dönmez")
	assert.Equal(t, mm["ismail@pardus.org.tr"].Email, "")

	assert.Equal(t, mm["linux@horizon.com"].Name, "anonymous")
	assert.Equal(t, mm["linux@horizon.com"].Email, "")

	assert.Equal(t, mm["y0netan1@dragonflybsd.org"].Name, "YONETANI Tomokazu")
	assert.Equal(t, mm["y0netan1@dragonflybsd.org"].Email, "y0n3t4n1@gmail.com")

	assert.Equal(t, mm["qhwt+git@les.ath.cx"].Name, "YONETANI Tomokazu")
	assert.Equal(t, mm["qhwt+git@les.ath.cx"].Email, "y0n3t4n1@gmail.com")

	assert.Equal(t, mm["vnwildman@gmail.com"].Name, "Trần Ngọc Quân")
	assert.Equal(t, mm["vnwildman@gmail.com"].Email, "vnwildman@gmail.com")

	assert.Equal(t, mm["Tran Ngoc Quan"].Name, "Trần Ngọc Quân")
	assert.Equal(t, mm["Tran Ngoc Quan"].Email, "vnwildman@gmail.com")

	assert.Equal(t, mm["nico@cam.org"].Name, "")
	assert.Equal(t, mm["nico@cam.org"].Email, "nico@fluxnic.net")
}
