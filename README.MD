## LANCEMENT DU SERVER
Dans un terminal wsl : "go run main.go"
Se connecter avec :
    compte admin
        identifiant : admin
        mot de passe : admin
    compte utilisateur
        créer un compte sur : http://localhost:8080/register
# BEYOU FORUM
Ce readme correspond à notre dernier projet de B1 informatique : FORUM

## OBJECTIFS GLOBAUX DU PROJET
L'objectif de ce projet est de réaliser un forum en GOLANG. Pour cela il faut créer un serveur web, créer et exploiter une base de donnée SQLite. Ce projet nous permetra aussi de développer nos compétences en web.

## LANGAGES / OUTILS UTILISÉS
Le language du back est le golang (gestion de BDD, serveur web), le language du front est le html + css + js. Afin de faire un maquetage nous avons utilisé figma (lien figma : https://www.figma.com/file/spOOs4CXnSn8tGD0RWWn3S/Bword-Studio?type=design).

## ORGANISATION DU PROJET
Afin de nous organiser pour menner à bien notre projet nous nous sommes inspiré de la méthode agile. Nous avons découpé notre projet en une liste de tâches et nous avons fonctionné en rush. Afin d'optimisé cette organisation nous avons utilisé trello (lien trello : https://trello.com/b/fraINIRc/forum-beyou-studio). Pour partager les fichier au sein du groupe nous avons choisis d'utiliser github.

## Organisation BDD
Nous avons esseyé d'optimiser notre BDD afin de gérer un maximum grace au sql.
![Alt text](./static/img/BDD.png)

# RÉALISATION DU PROJET
### Page home (http://localhost:8080/)
Cette page est la page d'acceuil de notre site. Elle permet d'informer l'utilisateur sur le fontionnement de notre forum ainsi que de l'informer sur l'utilisation des cookies. Sur cette page l'utilisateur voit aussi les 3 topics les plus upvoté.
### Page de Register (http://localhost:8080/register)
Cette page permet à l'utilisateur de créé un compte afin de pouvoir acceder à toutes les fonctionnalité du forum. Si le nom d'utilisateur ou l'email est déja utilisé une pop up d'erreur s'affiche. S'il coche rember me il reste connecté à l'infinit sinon est connecté durant 6h.
### Page de Login (http://localhost:8080/login)
Cette page permet à l'utilisateur de se connecter afin de pouvoir acceder à toutes les fonctionnalité du forum. S'il se trompe sur son identifiant ou sur son mot de passe une pop up s'affiche. S'il coche rember me il reste connecté à l'infinit sinon est connecté durant 6h.
### Page d'account (http://localhost:8080/account)
Sur cette page l'utilisateur peux changer toutes ses informations (photo de profile, nom, prénom, num d'utilisateur, bio, mot de passe). Pour changer sa photo de profile il faut cliquer sur cellle ci. L'utilisateur peut supprimer définitivement son compte ainsi que se déconnecter.
### Page Profile (http://localhost:8080/profile)
Ici l'utilisateur peut voir ses topics et ses posts
### Page d'exploration (http://localhost:8080/explore)
Cette page est la page d'exploration de notre site. L'utilisateur y retrouve tout les topics et à la possibilité de trier les topics, ou d'en rechercher un precis en fonction de leur nom ou alors par les tags que le topic contient. C'est aussi sur cette page que l'utilisateur peut créer un topic (s'il est login).
## Page Topic (http://localhost:8080/topic/ + "nom_de_topic")
On accède en cliquant sur "check topic".
Sur cette page l'utilisateur peux lire les différents post lié à un topic et / ou en rajouter un. Il peut aussi upvote ou follow ce topic. Si l'utilisateur clique sur la photo de profil d'un utilisatgeur il accedera à son profil.
## Page admin (http://localhost:8080/automod)
On accède à cette page quand on est connecté en admin et qu'on clique sur repport dans la navBar.
Sur les pages il est possible de repport un post ou un topic grace au petit drapeau.
Tous les repports apparaissent dans cette page, l'admin peux alors ban l'utilisateur (icone marteau), suprimer le post (icone poubelle), et delete le report (icone croix).
L'admin peut aussi ajouter des mots bannis ou en retirer.
Quand un utilisateur est bannis tout son contenu est supprimé de la base de données et l'utilisateur apparait dans les "banned user", l'admin peut ainsi l'unban.