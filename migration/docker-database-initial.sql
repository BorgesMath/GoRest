create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Isaac Newton', 'Isaac Newton foi um físico, matemático e astrônomo inglês que desenvolveu as leis
 do movimento e a lei da gravitação universal.'),
('Albert Einstein', 'Albert Einstein foi um físico teórico alemão, conhecido por desenvolver a
 teoria da relatividade, uma das duas pilares da física moderna.'),
('Nikola Tesla', 'Nikola Tesla foi um inventor, engenheiro elétrico e 
físico sérvio, conhecido por suas contribuições no desenvolvimento do sistema de corrente alternada (CA).'),
('Marie Curie', 'Marie Curie foi uma física e química polonesa, pioneira 
no estudo da radioatividade. Foi a primeira mulher a ganhar um Prêmio Nobel 
e a única pessoa a ganhar em duas áreas científicas diferentes.'),
('Galileo Galilei', 'Galileo Galilei foi um astrônomo, físico e engenheiro 
italiano, muitas vezes chamado de "pai da astronomia observacional", "pai da física moderna" e "pai da ciência".'),
('James Clerk Maxwell', 'James Clerk Maxwell foi um físico escocês, 
conhecido por formular a teoria clássica da radiação eletromagnética, 
unindo eletricidade, magnetismo e luz em um só fenômeno.'),
('Charles Darwin', 'Charles Darwin foi um naturalista britânico, cuja 
teoria da evolução por seleção natural é uma das bases da biologia moderna.'),
('Michael Faraday', 'Michael Faraday foi um físico e químico inglês, 
conhecido por suas descobertas na indução eletromagnética e nas leis da eletrolise.'),
('Johannes Kepler', 'Johannes Kepler foi um astrônomo, matemático e 
astrólogo alemão, conhecido por suas leis do movimento planetário que 
contribuíram para o desenvolvimento da física celeste.'),
('Stephen Hawking', 'Stephen Hawking foi um físico teórico, cosmólogo
 e autor britânico. Ele foi um dos cientistas mais conhecidos por seu trabalho em buracos negros e a relatividade.'),
('Richard Feynman', 'Richard Feynman foi um físico teórico americano, 
conhecido por suas contribuições para a mecânica quântica e por popularizar a ciência através de suas palestras e livros.');
