<?php

require "./vendor/autoload.php";

use Illuminate\Hashing\BcryptHasher;

$a = new BcryptHasher();

$result = $a->make("nipeharefa");

var_dump($result);
die();