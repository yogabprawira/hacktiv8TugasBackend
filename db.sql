CREATE DATABASE yogablog;

USE yogablog;

CREATE TABLE `user` (
    `uid` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `name` VARCHAR(64) NULL DEFAULT NULL,
    `password_hash` VARCHAR(128) NULL DEFAULT NULL,
    `session` VARCHAR(512) NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);

INSERT into `user` values (null, 'yoga', 'Yoga Budhi Prawira', '', '');

CREATE TABLE `post` (
    `post_id` INT NOT NULL AUTO_INCREMENT,
    `post_title` TEXT,
    `post_content` TEXT,
    `post_author` VARCHAR(64) NULL DEFAULT NULL,
    `time_published` VARCHAR(64) NULL DEFAULT NULL,
    `is_published` INTEGER,
    PRIMARY KEY (`post_id`)
);

INSERT into `post` values (null, 'KIM JI-YOUNG, BORN 1982 (2019)', 'Mengadaptasi novel berjudul sama karya Cho Nam-joo, dalam presentasinya soal kesetaraan gender, Kim Ji-young, Born 1982 mengangkat salah satu permasalahan paling mendasar yang telah cukup sering diterjemahkan ke layar lebar, pun tiada modifikasi dari teknik penceritaan maupun departemen artistik. Terkesan biasa, kecuali anda mengetahui kondisi dinamika gender di Korea Selatan termasuk ragam kontroversi yang mengiringi novel serta perilisan filmnya.

Saya tidak perlu menjabarkan detail mengenai paham patriarki dan tindak seksisme yang mengakar kuat di sana. Cukup simak segelintir kasus terkait Kim Ji-young, Born 1982 berikut. Pasca mengaku membaca novelnya, Irene, anggota girl group Red Velvet, jadi bulan-bulanan warganet, menerima banyak ujaran kebencian, pun foto-foto dan merchandise yang menampilkan sosoknya dibakar. Sedangkan akun media sosial Jung Yu-mi dibanjiri hujatan setelah diumumkan bakal berperan di film ini.

Berdasarkan kondisi tersebut, Korea Selatan belum membutuhkan tontonan feminisme unik. Mereka butuh kegamblangan. Prianya butuh ditampar, wanitanya butuh dibangkitkan. Seperti judulnya, film ini menuturkan kisah hidup Kim Ji-young (Jung Yu-mi), yang setelah pernikahannya dengan Jung Dae-hyun (Gong Yoo) dikaruniai seorang puteri, terpaksa berhenti bekerja demi mengurus anak sebagai ibu rumah tangga. Sekilas segalanya normal, hingga Ji-young mulai bersikap aneh, bersikap layaknya orang lain. Kadang menjadi ibunya, di lain waktu menjadi kakaknya, pernah juga bicara bak mendiang neneknya.

“Mungkin dia kesurupan”, ujar seorang rekan kantor Dae-hyun, menunjukkan betapa rendahnya kesadaran masyarakat terkait kesehatan mental. Rutinitas Ji-young bukan saja melelahkan, pula repetitif. Mengurus pekerjaan rumah, menjaga anak, tanpa hiburan, tanpa kawan kecuali keluarga atau mantan rekan kerja yang jumlah kunjungannya bisa dihitung jari. Ji-young kebosanan, tapi bukan hanya itu penyebab mentalnya terguncang. Tidak pula filmnya memandang sebelah mata status ibu rumah tangga.

Poin utama dari naskah tulisan sang sutradara Kim Do-young bersama Yoo Young-A (Miracle in Cell No. 7, On Your Wedding Day, My Annoying Brother) adalah bahwa sang protagonis urung berkesempatan menentukan jalannya sendiri. Sepanjang hidup ia selalu terkekang, diatur oleh stigma-stigma terhadap wanita. Ketimbang bekerja, wanita lebih baik mengurus anak di rumah. Wanita semestinya dipekerjakan di dapur saja seperti pembantu. Wanita sebaiknya begini, wanita tidak boleh begitu. Tapi tak satu pun berkata, “Wanita seharusnya berhak menentukan hidupnya sendiri”.

Penonton dibawa melihat bagaimana di semua sisi kehidupannya, Ji-young pernah jadi korban seksisme. Di rumah, sang ayah lebih memperhatikan adik laki-lakinya. Di kantor, karirnya terhambat akibat kekhawatiran atasan bahwa ia takkan bisa berkontribusi maksimal untuk jangka panjang sebab kelak bakal dipusingkan urusan anak. Ibu mertuanya menyuruh Ji-young bekerja di dapur seharian, lalu menghadiahinya sebuah celemek. Para pria menganggap kehidupannya enak karena bisa bersantai minum kopi memakai penghasilan suami. Pendidikan tingginya terbuang percuma, sebagaimana jadi bahan bercandaan para ibu yang bernasib serupa dengannya. “Aku berkuliah teknik agar bisa mengajari anakku matematika” atau “Aku berkuliah di jurusan akting sebagai bekal membacakan dongeng pengantar tidur”, begitu mereka berkelakar.

Semua dilemparkan secara gamblang, nihil kesubtilan, yang mana biasanya merupakan titik lemah, tapi kembali lagi, kita harus melihat kondisi. Pada tempat di mana WC umum masih jadi tempat menakutkan selaku lahan pria mesum merekam para wanita dan korban pelecehan disalahkan akibat pakaian yang dikenakan, tuturan subtil bukanlah prioritas. Kim Ji-young, Born 1982 mengemban tugas membuka mata publik selebar mungkin dengan baik.

Apalagi kesan terlampau gamblang itu mampu dibayar lunas oleh keberhasilannya menghantarkan emosi. Akting Jung Yu-mi berjasa besar di sini. Darinya, apa yang tampak bukan lagi “sebatas” wanita yang sedih. Hatinya bukan lagi terluka, tapi berlubang. Lubang menganga itu menyisakan kehampaan yang terasa nyata dari sorot mata Yu-mi. Sorot mata wanita yang tidak lagi tahu mesti berbuat apa. Begitu sakit, mungkin ia tak lagi merasakan apa pun. Kadang sosoknya bak seonggok boneka, yang tercipta dari manusia yang dilucuti haknya, sampai tiba di titik mulai percaya bahwa mungkin ia memang tak bernyawa apalagi berdaya.

Memerankan Mi-sook, ibunda Ji-young, Kim Mi-kyung tidak kalah hebat. Puncaknya adalah momen kala Mi-sook akhirnya tahu seberapa parah kondisi mental sang puteri. Lagi-lagi penyutradaraan Do-young enggan menyisakan ruang bagi kesubtilan dan memang tidak wajib. Selepas puluhan menit menyesakkan saat menyaksikan ketidakmampuan Ji-young menyuarakan isi hati, momen ini bagai katarsis di mana semua rasa tumpah bersama air mata.

Menarik disimak bagaimana film ini menggambarkan para pria. Beberapa memang bajingan sejati, sedangkan sisanya, terutama Dae-hyun, merupakan sosok baik hati yang turut serta menegakkan patriarki tanpa disadari. Dae-hyun sepenuh hati ingin membantu sang istri, tetapi kultur seksisme sayangnya telah sedemikian mendarah daging. Proses pembelajaran Dae-hyun jadi satu poin yang luput digarap mendalam, membuat konklusi penuh harap selepas dua jam kelam jadi kurang maksimal akibat kesan buru-buru.

Kim Ji-young, Born 1982, baik versi novel maupun film, mungkin takkan seketika mengubah Korea, apalagi dunia. Tapi cukup sebagai batu pijakan, penyulut pergerakan-pergerakan, yang diharapkan bakal menghalangi adanya Kim Ji-young lain, alias wanita-wanita yang akibat ketidakadilan, kehilangan sinarnya. We don’t wanna live in a world where the stars don’t shine.', 'yoga', '18 November 2019', 1);


INSERT into `post` values (null, 'FORD V FERRARI (2019)', 'Pernah menggarap western (3:10 to Yuma) dan film superhero bergaya western (Logan), wajar saat Ford v Ferrari garapan sutradara James Mangold memancarkan nuansa serupa. Mengedepankan dua protagonis dengan cowboy attitude (bahkan salah satunya mengenakan topi koboi), ini bak kisah koboi yang lebih modern, di mana alih-alih seekor kuda, mobil balap Ford GT40 jadi tunggangannya.

Ketika Henry Ford II (Tracy Letts) kebakaran jenggot akibat penjualan mobilnya menurun, Wakil Presiden Ford, Lee Iacocca (Jon Bernthal), mengusulkan pada sang CEO agar perusahaan mereka mengikuti jejak Ferrari berlomba di ajang “24 Hours of Le Mans”. Menurut Iacocca, keberhasilan pabrikan mobil milik Enzo Ferrari (Remo Girone) itu menyabet gelar juara secara beruntun membuatnya tampak superior hingga digandrungi publik.

Awalnya Ford berniat membeli Ferrari yang nyaris bangkrut, namun di “tikungan akhir”, Fiat menyalip mereka. Tersinggung oleh penolakan serta hinaan Enzo, Henry mengubah rencana. Ditunjuklah Carroll Shelby (Matt Damon) guna membuatkan Ford sebuah mobil balap yang lebih cepat dari Ferrari. Shelby, yang pernah menjuarai Le Mans 1959 sebelum pensiun akibat gangguan jantung, turut mengajak rekannya, Ken Miles (Christian Bale), pembalap bertalenta luar biasa yang kerap dicap buruk akibat perangai urakannya.

Shelby boleh membuka cerita, dan Matt Damon sendiri tampil solid tanpa perlu terkesan “showy” layaknya setir yang mengontrol laju filmnya, tapi mesin penggerak Ford v Ferrari adalah Miles. Di hadapan banyak orang, gestur, ekspresi, sampai cara bicara Bale menggambarkan betul bagaimana Miles memposisikan dirinya di atas lawan bicara. Apalagi kala mesti beradu melawan Leo Beebe (Josh Lucas), salah satu eksekutif Ford yang akan membuat penonton ingin melayangkan bogem mentah ke arah senyum kesombongan pihak korporat yang senantiasa ia pasang di wajahnya.

Sebaliknya, di ruang intim, baik di tengah kesendirian, di balik kemudi mobil, maupun di samping keluarganya, Miles adalah manusia biasa, yang berperasaan dan kerap menunjukkan kerapuhan. Dinamika keluarga Miles justru merupakan roh filmnya. Caitriona Balfe memerankan Mollie, istri Miles yang suportif namun tak pasif, dan berani bersuara kala sang suami melakukan kesalahan. Sedangkan Noah Jupe adalah Peter, putera Miles yang amat mengagumi sang ayah.

Naskah buatan kakak beradik Jez Butterworth dan John-Henry Butterworth (Edge of Tomorrow, Get on Up) bersama Jason Keller (Mirror Mirror, Escape Plan) paham betul bahwa drama olahraga terbaik selalu soal sisi personal pelakunya. Alhasil, momen emosional Ford v Ferrari selalu soal keluarga Miles. Ketika mobil Miles meledak di sesi latihan, kita diajak merasakan teror mencekam seorang anak yang menyaksikan ayahnya mendekati maut. Pun bukan diskusi teknis yang dipakai untuk menjabarkan luar biasa panjang, lama, nan menantangnya “24 Hours of Le Mans”, melainkan obrolan hati ke hati Miles dan Peter.

Di lintasan balap, giliran James Mangold unjuk gigi. Semua balapan digarap maksimal, tidak ada yang sekadar numpang lewat. Dari perlombaan Daytona sampai Le Mans punya ketegangan sekaligus euforia masing-masing. Dibantu sinematografer langganannya, Phedon Papamichael, ditambah penyuntingan cekatan, Mangold membangun intensitas melalui penempatan kamera yang mencakup seluruh sisi. Kita tahu kondisi di dalam mobil termasuk ekspresi Miles, sudut-sudut lintasan, pula bagaimana mobil melaju di sana. Selaku latar, musik gubahan Marco Beltrami (The Hurt Locker, A Quiet Place) memadukan beragam bentuk, dari sentuhan jazz (Ferrari Factory, Photos to Fiat) hingga rock pemacu adrenalin (Le Mans 66, Willow Sprints).

Seru, menegangkan, dan emosional, Ford v Ferrari bukan tentang kedigdayaan dua pabrikan mobil tersebut. Bahkan hingga akhir, filmnya tetap konsisten melontarkan kritik terhadap pihak korporat yang melakukan segala cara demi keuntungan sendiri. Sekali lagi, drama olahraga terbaik selalu bicara seputar sisi personal pelakunya, dan Ford v Ferrari berhasil melakukan itu, menyoroti perjuangan dua koboi lintasan balap, menjadikannya salah satu yang terbaik dalam beberapa waktu terakhir.', 'yoga', '18 November 2019', 1);


CREATE TABLE `message` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `full_name` TEXT,
    `email` TEXT,
    `message` TEXT,
    PRIMARY KEY (`id`)
);
