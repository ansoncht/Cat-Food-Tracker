package com.ansoncht.catfoodtracker.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document("users")
public class UserDAO {
    @Id private String id;
    @NotBlank private String firstName;
    @NotBlank private String lastName;
    @NotBlank private String username;
    @NotBlank @Email private String email;
    @NotBlank private String password;

    private UserDAO() {}

    private UserDAO(
            String firstName, String lastName, String username, String email, String password) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.username = username;
        this.email = email;
        this.password = password;
    }

    public static UserDAO createUserDAO(UserDTO.SignUpRequest userDTO) {
        return new UserDAO(
                userDTO.getFirstName(),
                userDTO.getLastName(),
                userDTO.getUsername(),
                userDTO.getEmail(),
                userDTO.getPassword());
    }

    public static UserDAO createUserDAO(
            String firstName, String lastName, String username, String email, String password) {
        return new UserDAO(firstName, lastName, username, email, password);
    }

    public String getId() {
        return this.id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getFirstName() {
        return this.firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return this.lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getUsername() {
        return this.username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getEmail() {
        return this.email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return this.password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}
