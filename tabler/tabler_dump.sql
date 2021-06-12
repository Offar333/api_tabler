-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: tabler_db
-- ------------------------------------------------------
-- Server version	8.0.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `formatos`
--

DROP TABLE IF EXISTS `formatos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `formatos` (
  `FORM_ID` int NOT NULL AUTO_INCREMENT,
  `FORM_DESC` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`FORM_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `formatos`
--

LOCK TABLES `formatos` WRITE;
/*!40000 ALTER TABLE `formatos` DISABLE KEYS */;
INSERT INTO `formatos` VALUES (1,'A Song of Ice and Fire'),(2,'Battle Century G'),(3,'BattleTech'),(4,'BattleTech 4th Edition'),(5,'BattleTech: Time of War'),(6,'BEER Engine!'),(7,'Better Angels'),(8,'Beyond the Wall and Other Adventures'),(9,'Black-Seven'),(10,'Blightburg'),(11,'Brigandyne'),(12,'Brigandyne (French)'),(13,'Bulletproof Blues'),(14,'Buried Treasures'),(15,'Burn Bryte'),(16,'Butterfly Aspect'),(17,'C.O.P.S'),(18,'Cahiers du Vastemonde'),(19,'Call of Cthulhu 5th Edition'),(20,'Call of Cthulhu 6th Edition'),(21,'Call of Cthulhu 6th Edition (French)'),(22,'Call of Cthulhu 7th Edition'),(23,'Call of Cthulhu: Dark Ages'),(24,'Capes, Cowls & Villains Foul'),(25,'Capharnaum'),(26,'Capitaine Vaudou'),(27,'Carbon 2185'),(28,'Castles & Crusades'),(29,'CdB Engine'),(30,'Celerity'),(31,'Celerity Lite'),(32,'Cepheus Engine'),(33,'Champions Now'),(34,'Channel Fear'),(35,'Channel Fear (French)'),(36,'Chaos 6010'),(37,'Chaos System'),(38,'Chaosium'),(39,'Chronicle System'),(40,'Chroniques Galactiques (French)'),(41,'Chroniques Oubliées'),(42,'Chroniques Oubliees Contemporain'),(43,'Chroniques Oubliées Fantasy (French)'),(44,'Clockwork System'),(45,'Conan 2nd Edition'),(46,'Continuum'),(47,'COPS'),(48,'Cortex Plus'),(49,'Cortex Plus Heroic'),(50,'Cortex System'),(51,'Creation of Crônicas'),(52,'Criancas Enxeridas'),(53,'Crônicas'),(54,'Cthulhu D100'),(55,'Cthulhu Dark'),(56,'Cthulhu Tech'),(57,'Cuentos de Ánimas'),(58,'Cultos Innombrables (Spanish)'),(59,'Cyberpunk Red'),(60,'Cyberpunk Red Jumpstart'),(61,'Cypher System'),(62,'D&D 3.5E'),(63,'D&D 3.5E Adventure Pack'),(64,'D&D 3.5E The Black Company'),(65,'D&D 4E'),(66,'D&D 5E (Community Contributed)'),(67,'D&D 5E (Community, Chinese)'),(68,'D&D 5E (Darker Dungeons)'),(69,'D&D 5E (Scrolls)'),(70,'D&D 5E (Shaped)'),(71,'D&D 5E Adventures in Middle Earth'),(72,'D&D Retroclone'),(73,'d20 Modern'),(74,'d20 System'),(75,'D6'),(76,'D6 Adventure'),(77,'D6 Fantasy'),(78,'D6 Space'),(79,'Dangerous Journeys'),(80,'Dangerous Journeys: Mythus'),(81,'Das Schwarze Auge'),(82,'Das Schwarze Auge 4th Edition (German)'),(83,'Das Schwarze Auge 5th Edition (German)'),(84,'Das Schwarze Auge 5th Edition (Tsa Edition)'),(85,'Das Schwarze Auge Ilaris'),(86,'Deadlands Classic'),(87,'Deadlands Classic Anniversary - Improved'),(88,'Deadlands: Classic'),(89,'Deadlands: Classic Anniversary'),(90,'Degenesis'),(91,'Degenesis (Advanced)'),(92,'Delta Green'),(93,'Demon Gate'),(94,'Demon Gate Scripted'),(95,'Descent into Midnight'),(96,'Devâstra'),(97,'Digital Shades'),(98,'DK System'),(99,'Dogs in the Vineyard'),(100,'Dominion'),(101,'Dragon De Poche'),(102,'Dragon De Poche 2 (French)'),(103,'Dragon Warriors'),(104,'Dragon Warriors Revised'),(105,'Dragoncry'),(106,'DragonQuest'),(107,'Drakar och Demoner 6.0 (Swedish)'),(108,'Drakar och Demoner 91'),(109,'Dramasystem'),(110,'Dreamchaser'),(111,'Dungeon Crawl Classics'),(112,'Dungeon Crawl Classics (Tabbed)'),(113,'Dungeon Squad'),(114,'Dungeons & Dragons 3.5 Edition'),(115,'Dungeons & Dragons 4th Edition'),(116,'Dungeons & Dragons 5th Edition'),(117,'Eirendor'),(118,'El Resurgir del Dragón'),(119,'Esper Genesis'),(120,'Firefly'),(121,'Héros & Dragons'),(122,'Iron Kingdoms'),(123,'Kobolds Ate My Baby'),(124,'Lamentations of the Flame Princess'),(125,'Leverage'),(126,'LME'),(127,'Magic World'),(128,'Marvel Heroic Roleplaying'),(129,'Mass Effect'),(130,'MechWarrior Destiny'),(131,'Mutant (1984)'),(132,'Mutant Crawl Classics'),(133,'Naruto'),(134,'Nephilim'),(135,'No Thank You, Evil'),(136,'Numenera'),(137,'RuneQuest 3rd Edition'),(138,'RuneQuest 6th Edition'),(139,'RuneQuest 6th Edition (Spanish)'),(140,'RuneQuest Glorantha'),(141,'RuneQuest: 6th Edition (French)'),(142,'Serenity'),(143,'Song of Ice and Fire Roleplaying'),(144,'Star Wars 30th Anniversary Edition'),(145,'Star Wars 5e'),(146,'Star Wars Revised'),(147,'Star Wars Saga Edition'),(148,'Star Wars: (D6) West End Games'),(149,'Stormbringer'),(150,'Tephra'),(151,'Terre 2'),(152,'The Laundry'),(153,'The Strange (Multiple Recursion)'),(154,'Vitu Lindao');
/*!40000 ALTER TABLE `formatos` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mesa`
--

DROP TABLE IF EXISTS `mesa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mesa` (
  `ID_MESA` int NOT NULL AUTO_INCREMENT,
  `ADM_MESA` varchar(15) DEFAULT NULL,
  `TITULO_MESA` varchar(45) NOT NULL,
  `DESC_MESA` varchar(100) NOT NULL,
  `QTDEJOG_MESA` varchar(20) DEFAULT NULL,
  `FORMA_MESA` varchar(50) NOT NULL,
  `STATUS_MESA` varchar(15) NOT NULL,
  `DATCRIAC_MESA` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `LVLINIC_MESA` varchar(3) DEFAULT NULL,
  `EXPJOGO_MESA` varchar(45) DEFAULT NULL,
  `LINK_CHAT` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`ID_MESA`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mesa`
--

LOCK TABLES `mesa` WRITE;
/*!40000 ALTER TABLE `mesa` DISABLE KEYS */;
INSERT INTO `mesa` VALUES (1,'Offar','teste de alteracao2','Aqui é informado a descrição da mesa. Essa é uma descrição de teste','3','D&D 5e','1','2021-04-20 00:07:41','1','teste','t'),(2,'Offar','teste de alteracao','desc alterada','3','D&D 5e','1','2020-09-01 17:23:59','1','teste','t'),(3,'Vitu','TESTANDO 123','DVD EH UM HOME SEXUAL','8','DND','1','2020-09-01 17:23:59','1','teste','t'),(4,'XANDA','XANDAO','XANDAO GOSTA DE OFERECER O 2 DE AGOSTO','2','TESTE','1','2020-09-01 17:21:31','1','teste','t'),(5,'NAOSE','OQ TA','ACONTECENO NESSA PORRA','1','AJUDA NOIS PLZ','1','2020-09-01 17:21:31','1','teste','t'),(6,'Virgolino','Descent into Avernus','Campanha descent into avernus','5','D&D5e','1','2021-04-19 23:30:38','1','teste','t'),(7,'AdmTeste1','Mesa1','Descricao Mesa 1','2','Formato Mesa 1','2','2020-09-01 17:21:31','1','teste','t'),(8,'AdmTeste2','Mesa2','Descricao Mesa 2','2','Formato Mesa 2','2','2020-09-01 17:21:31','1','teste','t'),(9,'AdmTeste3','Mesa3','Descricao Mesa 3','2','Formato Mesa 3','2','2020-09-01 17:21:31','1','teste','t'),(10,'AdmTeste4','Mesa4','Descricao Mesa 4','2','Formato Mesa 4','2','2020-09-01 17:21:31','1','teste','t'),(11,'AdmTeste5','Mesa5','Descricao Mesa 5','2','Formato Mesa 5','2','2020-09-01 17:21:31','1','teste','t'),(12,'AdmTeste6','Mesa6','Descricao Mesa 6','2','Formato Mesa 6','2','2020-09-01 17:21:31','1','teste','t'),(13,'AdmTeste7','Mesa7','Descricao Mesa 7','2','Formato Mesa 7','2','2020-09-01 17:21:31','1','teste','t'),(14,'AdmTeste8','Mesa8','Descricao Mesa 8','2','Formato Mesa 8','2','2020-09-01 17:21:31','1','teste','t'),(15,'AdmTeste9','Mesa9','Descricao Mesa 9','2','Formato Mesa 9','2','2020-09-01 17:21:31','1','teste','t'),(16,'Neandertal','australopiteco','teste de australopiteco','1','Formato Primal','1','2020-09-01 17:21:31','1','teste','t'),(19,'1','Ilha do Molusco','Procura-se penis de 20 cm','','Call of Cthulhu 7th Edition','0','2020-09-01 17:21:31','1','Sênior','t'),(20,'1','Segredo Moluscao','Procura-se Rafael','','Cthulhu Dark','0','2020-09-01 17:21:31','1','Sênior','t'),(21,'1','Joao bolao','meu pau na tua mao','','Cypher System','0','2020-09-01 17:21:31','1','Veterano','t'),(22,'Offar','testeqtdejog','mesa de testeqtdejog','3','D&D 5e','1','2020-09-01 17:21:31','1','teste','t'),(23,'1','Agoravai mano','juro po se','6','Crônicas','0','2020-09-01 17:21:31','4','Sênior','t'),(24,'Offar','Forja Gozada','Procura-se anões musculosos','5','Iron Kingdoms','0','2020-09-01 17:21:31','1','Pleno','t'),(25,'Offar','teste mesa','asdasda','5','Cthulhu Dark','0','2020-09-01 17:21:31','1','Júnior','t'),(26,'Offar','testererer','sdasdasd','5','D&D 5E Adventures in Middle Earth','0','2020-09-01 17:21:31','1','Júnior','t'),(27,'Offar','Rafael Teste','Teste Rafael','8','D&D Retroclone','0','2020-09-01 15:33:48','3','Sênior','https://paunocudequemtalendo.com.br/fdKKDS0dDS'),(31,'Offar','Tesderedesu Rafaele','Faraele','7','D6','0','2020-09-01 16:48:57','4','Pleno','https://xvideos.com/Comiairmadomeuamigoconfiraareacaodele'),(33,'Offar','testeqtdejog','mesa de testeqtdejog','3','D&D 5e','1','2020-09-01 17:00:20','1','teste','qerqwerqwer'),(34,'Offar','Vitória ória','Vitória ória','8','Cthulhu Tech','0','2020-09-01 17:11:57','2','Pleno','https://birdmemes.com/jaaaj'),(35,'Offar','Gaudales','um','12','d20 System','0','2020-09-01 20:16:24','10','Pleno','https://puts.com'),(36,'Offar','teru','gcyvh','5','D&D 5E Adventures in Middle Earth','0','2020-09-01 21:15:13','1','Júnior','hhy'),(37,'Offar','Teste de Criação','Ta funfando','10','D6 Fantasy','0','2020-09-10 18:55:34','1','Pleno','https://comiocudequemtalendo.com'),(38,'UserTeste','TesteRafa','Yooow','2','D&D 5E Adventures in Middle Earth','0','2021-04-11 23:24:19','1','Júnior','https://agenciaolcan.com.br'),(39,'UserTeste','Teste','asdasd','5','D&D 5E Adventures in Middle Earth','0','2021-04-11 23:28:45','1','Júnior','asdasd');
/*!40000 ALTER TABLE `mesa` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mesa_jogadores`
--

DROP TABLE IF EXISTS `mesa_jogadores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mesa_jogadores` (
  `ID_JOGA` int NOT NULL AUTO_INCREMENT,
  `ID_MESA` int NOT NULL,
  `ID_USUAR` int NOT NULL,
  `MESTRE_JOGA` tinyint(1) NOT NULL,
  `FICHA_JOGA` varchar(254) DEFAULT NULL,
  `NOMECHAR_JOGA` varchar(45) DEFAULT NULL,
  `CLASSECHAR_JOGA` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`ID_JOGA`),
  KEY `ID_USUAR` (`ID_USUAR`),
  KEY `ID_MESA` (`ID_MESA`),
  CONSTRAINT `mesa_jogadores_ibfk_1` FOREIGN KEY (`ID_USUAR`) REFERENCES `usuario` (`ID_USUAR`),
  CONSTRAINT `mesa_jogadores_ibfk_2` FOREIGN KEY (`ID_MESA`) REFERENCES `mesa` (`ID_MESA`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mesa_jogadores`
--

LOCK TABLES `mesa_jogadores` WRITE;
/*!40000 ALTER TABLE `mesa_jogadores` DISABLE KEYS */;
INSERT INTO `mesa_jogadores` VALUES (1,1,1,0,'c:/brabo','teste','teste'),(2,2,1,0,'C:/taaquiaficha','teste1','teste1'),(3,2,3,1,'C:/taaquiaficha','teste2','teste2'),(4,2,2,0,'C:/taaquiaficha','teste3','teste3'),(5,1,2,0,'C:/SheetPath','Fufusca','Guerreiro'),(6,3,40,0,'C:/SheetPath','teste4','teste4'),(7,1,3,0,'C:/ehissoae','Gilson','Mago'),(9,3,2,0,'teste','teste6','teste6'),(10,3,3,0,'teste','teste7','teste7'),(11,3,41,1,'teste','teste8','teste8'),(12,10,1,1,'teste','teste9','teste9'),(14,5,1,0,'CharSheet/Offar/Osmar','Osmar','Guerreiro'),(15,7,1,0,'CharSheet/Offar/teste','teste','aaaa'),(18,35,1,1,'Mestre','Mestre','Mestre'),(19,34,1,0,'CharSheet/Offar/te','te','ste'),(20,6,1,0,'CharSheet/Offar/jgf','jgf',''),(22,21,1,1,'Mestre','Mestre','Mestre'),(23,25,1,1,'Mestre','Mestre','Mestre'),(24,36,1,0,'CharSheet/Offar/a','a','r'),(25,31,1,1,'Mestre','Mestre','Mestre'),(26,16,1,0,'CharSheet/Offar/teste','teste','asdasd'),(27,26,1,0,'CharSheet/Offar/Farael','Farael','gostoso'),(28,24,1,0,'CharSheet/Offar/Joania','Joania','sfada'),(29,23,1,0,'CharSheet/Offar/atwhdb','atwhdb','sesbsh'),(30,19,1,0,'CharSheet/Offar/sdasda','sdasda','asdasd'),(31,22,1,0,'CharSheet/Offar/teste','teste','hbyhhg'),(32,20,1,0,'CharSheet/Offar/Teste','Teste','ajbdahbd'),(33,11,1,0,'CharSheet/Offar/aaaa','aaaa','sssss'),(34,12,1,0,'CharSheet/Offar/nsdjjds','nsdjjds','snsnnss'),(35,13,1,0,'CharSheet/Offar/jaaaj','jaaaj','jooooj'),(36,14,1,1,'Mestre','Mestre','Mestre'),(37,37,1,0,'CharSheet/Offar/Josivaldo','Josivaldo','Archero'),(38,1,68,1,'Mestre','Mestre','Mestre'),(39,6,68,0,'CharSheet/UserTeste/Josefo','Josefo','1234'),(40,3,68,0,'CharSheet/UserTeste/Josefoe','Josefoe','1234'),(41,4,68,1,'Mestre','Mestre','Mestre'),(42,5,68,1,'Mestre','Mestre','Mestre'),(43,36,68,0,'CharSheet/UserTeste/joao ricardo','joao ricardo','Musico'),(44,37,68,1,'Mestre','Mestre','Mestre'),(45,35,68,0,'CharSheet/UserTeste/joooa','joooa','ssds'),(46,12,68,1,'Mestre','Mestre','Mestre'),(47,27,69,1,'Mestre','Mestre','Mestre'),(48,34,69,1,'Mestre','Mestre','Mestre'),(49,27,1,0,'CharSheet/Offar/globonildo','globonildo','clerigo do pao de alho'),(50,4,1,0,'CharSheet/Offar/Dougloncio','Dougloncio','cuzao'),(51,4,69,0,'CharSheet/Vitinho do Sus/Viiv sus','Viiv sus','Rafao'),(52,6,69,1,'Mestre','Mestre','Mestre'),(53,7,68,0,'CharSheet/UserTeste/Testonildo','Testonildo','guerreiro'),(54,8,68,1,'Mestre','Mestre','Mestre'),(55,24,68,1,'Mestre','Mestre','Mestre'),(56,23,68,1,'Mestre','Mestre','Mestre'),(57,38,68,1,'Mestre','Mestre','Mestre'),(58,19,68,1,'Mestre','Mestre','Mestre'),(59,15,68,1,'Mestre','Mestre','Mestre'),(60,13,68,1,'Mestre','Mestre','Mestre'),(61,39,68,1,'Mestre','Mestre','Mestre');
/*!40000 ALTER TABLE `mesa_jogadores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessoes_ativas`
--

DROP TABLE IF EXISTS `sessoes_ativas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sessoes_ativas` (
  `ID_SESS` int NOT NULL AUTO_INCREMENT,
  `ID_USUAR` int DEFAULT NULL,
  `DATA_SESS` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID_SESS`),
  KEY `ID_USUAR` (`ID_USUAR`),
  CONSTRAINT `sessoes_ativas_ibfk_1` FOREIGN KEY (`ID_USUAR`) REFERENCES `usuario` (`ID_USUAR`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessoes_ativas`
--

LOCK TABLES `sessoes_ativas` WRITE;
/*!40000 ALTER TABLE `sessoes_ativas` DISABLE KEYS */;
/*!40000 ALTER TABLE `sessoes_ativas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuario`
--

DROP TABLE IF EXISTS `usuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuario` (
  `ID_USUAR` int NOT NULL AUTO_INCREMENT,
  `NOME_USUAR` varchar(30) NOT NULL,
  `APELIDO_USUAR` varchar(15) NOT NULL,
  `SENHA_USUAR` varchar(15) NOT NULL,
  `EMAIL_USUAR` varchar(30) NOT NULL,
  `AVATAR_USUAR` varchar(254) DEFAULT NULL,
  PRIMARY KEY (`ID_USUAR`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuario`
--

LOCK TABLES `usuario` WRITE;
/*!40000 ALTER TABLE `usuario` DISABLE KEYS */;
INSERT INTO `usuario` VALUES (1,'Rafael Vilas','Offar','123456','rafael.gostosa@.com','ProfilePics/Offar'),(2,'Vitor','Vitu','123456','testeteste','C:/Imagens/FotinhaManeira.jpg'),(3,'Clebin','Yrton','32546','email@alterado','D:/Imagens/FotinhaManeira.jpg'),(40,'Ronaldinho','Gaucho','1234','gr@gmail','C:/FotoPadrao'),(41,'Douglas Henrique','Tupete','123456','ajlshd@qwejhk','C:/FotoPadrao'),(42,'Johnny','Joooj','1234','Joooj@hotmail.com','C:/FotoPadrao'),(43,'tesdsd','23234','','sdd','C:/FotoPadrao'),(44,'Corno','Rafael','dfddf','fdsdf','C:/FotoPadrao'),(46,'Piroca','Murcha','piroca','piroquinhamurcha@hotmail.com','C:/FotoPadrao'),(48,'Jooj','','','','C:/FotoPadrao'),(49,'Johnny Carvalho','Carvalho JooJ','123456789','Jooj@Jooj.com.br','C:/FotoPadrao'),(54,'Ronaldão','Fenomeno','123456','teste@terra.com.br','C:/FotoPadrao'),(55,'teste','asdasd asd','asdfas','sdsd','C:/FotoPadrao'),(56,'Sou um usuário idiota','nossa','senha','vopo@oemail.aqui','C:/FotoPadrao'),(57,'Teste','joao','pablo','pablo','C:/FotoPadrao'),(58,'text','text','','text','C:/FotoPadrao'),(59,'carai','mermao','mermao','asa','C:/FotoPadrao'),(60,'Af','as','','aaaa','C:/FotoPadrao'),(61,'Vitor H N Piovezan','Vituh','vitor2142','vitor.piovezan7@gmail.com','C:/FotoPadrao'),(62,'joaozinhi','joana','','coalinha','C:/FotoPadrao'),(63,'joaozinhi','joanhinha','tefds','joaninha@joaninha','C:/FotoPadrao'),(64,'Vitoria','Viiiiiiiihx','1234','vi@vi.vi','C:/FotoPadrao'),(65,'Mika Ovileira','mika','','mika@mika.mika','C:/FotoPadrao'),(66,'Mika Ovileira','mikae','mika','mika@mika@mika.com','C:/FotoPadrao'),(67,'Daniela','daniela','123','rafaelvcarvalho@hotmail.com','C:/FotoPadrao'),(68,'UserTeste1','UserTeste','1234','user.teste','C:/FotoPadrao'),(69,'Vitor SUSS','Vitinho do Suss','vitu','vitu@vi.tu','C:/FotoPadrao');
/*!40000 ALTER TABLE `usuario` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-06-11 19:21:34
